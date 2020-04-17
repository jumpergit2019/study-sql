package study_gorm

import "github.com/jinzhu/gorm"

func ExampleDelete(db *gorm.DB) {
	////////////////////////////////////delete data
	db.Where("id = ?", 7).Delete(&Animal{})
	db.Where("id = ?", 6).Unscoped().Delete(&Animal{})

	//若是存在delete at , 那么使用 delete 执行的是软删除，要执行硬删除需要使用 unscoped delete
	//软删除了的数据 query 无法显现， 要显现需要在query中使用 unscoped

}
