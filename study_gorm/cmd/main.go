package main

import (
	"database/sql"
	"time"

	"github.com/jumpergit2019/study-sql/study_gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//warning: 目前 外键 存在bug, 无法设置。先跳过 Associations 部分
//warning: 目前 db.Transaction 存在bug, 无法处理 panic

func main() {
	db, err := gorm.Open("mysql", "jumper:J1u2m3p!e@r#@tcp(192.168.1.35:3306)/test?loc=Local&parseTime=true&charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//db.SetLogger(gorm.Logger{})
	db.LogMode(true)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)

	db.DropTableIfExists("animals")
	db.DropTableIfExists("owners")

	db.AutoMigrate(&study_gorm.Animal{})
	db.AutoMigrate(&study_gorm.Owner{})

	initData(db)

	study_gorm.ExampleCreate(db)
	study_gorm.ExampleQuery(db)
	study_gorm.ExampleUpdate(db)
	study_gorm.ExampleDelete(db)
}

func initData(db *gorm.DB) error {
	animals := []study_gorm.Animal{
		{
			Name: "wang",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 18,
			},
			OwnerId: 1,
		},
		{
			Name: "zhang",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 18,
			},
			OwnerId: 2,
		},
		{
			Name: "li",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 28,
			},
			OwnerId: 3,
		},
		{
			Name: "feng",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 38,
			},
			OwnerId: 4,
		},
		{
			Name: "huang",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 56,
			},
			OwnerId: 5,
		},
		{
			Name: "tan",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 35,
			},
			OwnerId: 6,
		},
		{
			Name: "xiang",
			Age: sql.NullInt32{
				Valid: true,
				Int32: 26,
			},
			OwnerId: 7,
		},
	}

	for _, animal := range animals {
		db.Create(&animal)
	}

	owners := []study_gorm.Owner{
		{
			Name: "niuniu",
		},
		{
			Name: "xiaoxiao",
		},
		{
			Name: "dong",
		},
		{
			Name: "sbzhu",
		},
		{
			Name: "ss",
		},
		{
			Name: "aa",
		},
	}
	for _, owner := range owners {
		db.Create(&owner)
	}

	return nil
}
