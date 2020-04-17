package study_gorm

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ExampleQuery(db *gorm.DB) {
	////////////////////////////////////query data
	var a Animal
	db.First(&a)
	fmt.Println("a: ", a)

	var a10 Animal
	db.First(&a10, 10)
	fmt.Println("a10: ", a10)

	var b Animal
	db.Last(&b)
	fmt.Println("b: ", b)

	var as []Animal
	db.Find(&as)
	fmt.Println("as: ", as)

	var aa Animal
	db.Where("id = ?", 16).First(&aa)
	fmt.Println(aa)

	var aas []Animal
	db.Where("id > ?", 10).Find(&aas)
	fmt.Println(aas)

	var abs []Animal
	names := []string{"wang", "ru", "zhang", "zhao"}
	db.Where("name in (?)", names).Find(&abs)
	fmt.Println(abs)

	var acs []Animal
	db.
		//Debug().
		Where("age > ?", 30).
		Where("name in (?)", names).
		Or("age < ?", 18).
		Find(&acs)
	fmt.Println("acs: ", acs)

	var ads []Animal
	db.
		//Debug().
		Where("age between ? and ?", 30, 60).
		Find(&ads)
	fmt.Println("ads: ", ads)

	var ae Animal
	db.
		Debug().
		Set("gorm:query_option", "for update").
		Where("id = ?", 1).
		Find(&ae)
	fmt.Println(ae)

	//需要order by才使用 first / last
	//不需要排序直接使用 find , 查找结果多个就使用 [], 一个就使用变量
	//使用 where 方式，更加清晰，容易修改， 尽量不使用其他方式

	var af Animal
	ra := db.
		Debug().
		Where(&Animal{Name: "llllllll"}).
		Attrs(&Animal{Age: struct {
			Int32 int32
			Valid bool
		}{Int32: 111, Valid: true}}).
		FirstOrInit(&af).RowsAffected

	fmt.Println(af, ra)

	if ra == 0 {
		db.Create(&af)
	}

	var ag Animal
	db.
		Debug().
		Where(&Animal{Name: "ffffff"}).
		Attrs(&Animal{Age: struct {
			Int32 int32
			Valid bool
		}{Int32: 999, Valid: true}}).
		FirstOrCreate(&ag)
	fmt.Println(ag)

	//firstorinit 用于判断是否存在该记录，若是存在，就获取内容，若是不存在就使用 attrs 创建内存对象，
	//之后可以判断 rows affected 来确定是否获取到值，没有则可以使用create 来插入数据。
	//上述 firstorinit + rows affected + create 等价于 firstorcreate
	//总结： 有就获取，没有就创建

	var ahs []Animal
	db.Debug().
		Where("age > ?", db.Table("animals").Select("avg(age)").SubQuery()).
		Find(&ahs)
	fmt.Println(ahs)

	//子查询需要使用 subquery, 使用方式如上

	var ais []Animal
	db.Debug().Select("name, age").Where("age > ?", 30).Find(&ais)
	fmt.Println(ais)

	var ajs []Animal
	var aks []Animal
	db.Debug().Where("name in (?)", []string{"wang", "zhang"}).Order("name desc").Find(&ajs).Order("age desc", true).Find(&aks)
	fmt.Println(ajs)
	fmt.Println(aks)

	//order("x asc, y desc") 等价于 order("x asc").order("y desc")
	//order("x asc").find(&a1).order("y desc", true).find(&a2)
	//相当于进行了两个查询，分别使用两个order

	var als []Animal
	//db.Limit(3).Find(&als)
	db.Debug().Offset(2).Limit(-1).Find(&als)
	fmt.Println(als)

	//发现 limit 可以单独使用， 但是 offset 不能单独使用，只能与 limit 一起使用
	//即便使用 offset(3).limit(-1) 也不行
	//也就是说不能实现 从某个记录开始往后面获取所有记录，发现mysql5.7 也不支持 limit 3,-1

	var cnt3 int32
	db.Model(&Animal{}).Where("age > ?", 30).Count(&cnt3)
	fmt.Println(cnt3)

	var ams []Animal
	db.Debug().Where("age > ?", 30).Find(&ams).Count(&cnt3)
	fmt.Println(ams)
	fmt.Println(cnt3)

	db.Debug().Table("animals").Select("count(distinct(name))").Count(&cnt3)
	fmt.Println(cnt3)

	//需要同时获取记录和数量，使用 find + count
	//只需要获取数量，使用 model + count
	//需要计算聚合函数，使用 table + select + count

	rows, err := db.Table("animals").Select("id, name, age").Rows()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t Animal
		err := rows.Scan(&t.ID, &t.Name, &t.Age)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(t)
	}

	fmt.Println("------------------")

	rows2, err := db.Table("animals").Select("name, sum(age) as ages").Group("name").Having("sum(age) > ?", 50).Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows2.Close()
	for rows2.Next() {
		var t Animal
		err := rows2.Scan(&t.Name, &t.Age)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(t)
	}

	var name string
	var age int32
	row := db.Table("animals").Where("name = ?", "wang").Select("name, age").Row() // (*sql.Row)
	row.Scan(&name, &age)
	fmt.Println(name, age)

	rows3, err := db.Model(&Animal{}).Where("name = ?", "wang").Select("id, name, age, owner_id").Rows() // (*sql.Rows, error)
	defer rows.Close()

	for rows3.Next() {
		var user Animal
		// ScanRows scan a row into user
		db.ScanRows(rows3, &user)
		fmt.Println(user)
	}

	//可以使用 db.ScanRows 将一行数据读入到一个结构体

	fmt.Println("------------------")

	type Tmp struct {
		Name string
		Ages int32
	}
	var tmps []Tmp
	var cnt int32
	db.Table("animals").Select("name, sum(age) as ages").Group("name").Having("sum(age) > ?", 50).Scan(&tmps).Count(&cnt)
	fmt.Println("tmps: ", tmps)
	fmt.Println("cnt: ", cnt)

	//一般查询也可以使用  table + select + rows, 相对于find/first 可以对每个数据进行进一步操作之后再存入内存（rows.next rows.scan)
	//分组使用 table + select + group + having + rows/scan

	///////////////////  https://gorm.io/docs/query.html joins
	rows4, err := db.Table("animals").Joins("join owners on animals.owner_id = owners.id").Where("owners.name = ?", "wang").Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows4.Close()

	ans := make([]Animal, 0)
	for rows4.Next() {
		var tmp Animal
		err := rows4.Scan(&tmp.ID, &tmp.Name, &tmp.Age, &tmp.OwnerId)
		if err != nil {
			fmt.Println(err)
			return
		}
		ans = append(ans, tmp)
	}

	fmt.Println(ans)

	var aos []Animal
	db.Table("animals").Joins("join owners on animals.owner_id = owners.id").Where("owners.name = ?", "wang").Scan(&aos)
	fmt.Println(aos)

	//join 可以多个join接连不断使用，表示多次join

	var ages []sql.NullInt32
	db.Table("animals").Where("name = ?", "wang").Pluck("age", &ages)
	fmt.Println(ages)

	type Result struct {
		Name string
		Age  int32
	}
	var result []Result
	//db.Table("animals").Select("name, age").Where("id in (?)", []int64{1, 5, 9, 12}).Scan(&result)
	db.Model(&Animal{}).Select("name, age").Where("id in (?)", []int64{1, 5, 9, 12}).Scan(&result)
	fmt.Println(result)
	//pluck 与 scan 类似，不过 pluck 只获取一列数据， scan 可以获取多列，获取哪几列由结构决定,但是需要使用 table/model 指明所查询对表名

}
