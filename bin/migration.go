package main

import (
	"github.com/k-kurikuri/last-didit-go/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main () {
	db, _ := gorm.Open("mysql", "root:root@/done_list?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DoneList{})
}
