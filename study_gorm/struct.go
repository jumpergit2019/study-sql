package study_gorm

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Animal struct {
	ID      int64
	Name    string        `gorm:"default:'galeone'"`
	Age     sql.NullInt32 `gorm:"default:18"`
	OwnerId int64
}

type Owner struct {
	ID   int64
	Name string
}

//func (animal *Animal) BeforeCreate(scope *gorm.Scope) error {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	scope.SetColumn("id", r.Int63n(100000))
//	return nil
//}

//type User struct {
//	ID      int64  `gorm:"primary_key;auto_increment:false"`
//	Name    string `gorm:"primary_key"`
//	Profile Profile
//}
//
//type Profile struct {
//	gorm.Model
//	UserID uint
//	Name   string
//	Age    int32
//}

func CreateAnimals(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
			return err
		}

		// return nil will commit
		return nil
	})
}
