package study_gorm

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ExampleUpdate(db *gorm.DB) {

	var aps Animal
	db.First(&aps)
	aps.Name = "dw"
	db.Debug().Save(&aps)

	db.Model(&Animal{}).Where("id = ?", 1).Update("name", "xw")
	db.Table("animals").Where("id = ?", 1).Update("name", "xw")
	db.Model(&Animal{}).Where("id = ?", 2).Updates(map[string]interface{}{"name": "xz", "age": sql.NullInt32{
		Int32: 99,
		Valid: true,
	}})

	//save 会将所有列进行设置
	//update 用于修改单列
	//updates 用于修改多列， 参数使用 map[string]interface{} 来指定， 若是使用结构体指定 是不会修改指定为零值的列的
	//这三个方法会调用 before update / after update

	rowCnt := db.Model(&Animal{}).Where("id = ?", 1).UpdateColumn("name", "wang2").RowsAffected
	fmt.Println(rowCnt)
	db.Model(&Animal{}).Where("id = ?", 2).UpdateColumns(map[string]interface{}{"name": "zhang2", "age": sql.NullInt32{
		Int32: 18,
		Valid: true,
	}})

	//update column 与 update 相似，但是不会调用 before update / after update
	//update columns 与 updates 相似，但是不会调用 before update / after update
	//使用 rows affected 获取实际修改行数
}
