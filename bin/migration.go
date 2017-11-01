package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k-kurikuri/gogo-done/app/models"
)

func main() {
	db, _ := gorm.Open("mysql", "root:root@/done_list?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.DoneList{},
		&models.Category{},
		&models.DoneListCategory{},
		&models.User{},
	)
}
