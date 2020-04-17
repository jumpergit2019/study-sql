package study_gorm

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ExampleCreate(db *gorm.DB) {

	////////////////////////////////////create data
	animal := Animal{
		Name: "",
		Age: sql.NullInt32{
			Int32: 0,
			Valid: false,
		},
	}
	fmt.Println("is new: ", db.NewRecord(&animal))
	db.Create(&animal)
	fmt.Println("is new: ", db.NewRecord(&animal))

}
