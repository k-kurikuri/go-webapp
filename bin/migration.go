package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/models"
)

func main() {
	db := db.Connection()

	defer db.Close()

	db.DropTableIfExists(
		&models.DoneList{},
		&models.Category{},
		&models.DoneListCategory{},
		&models.User{},
	)

	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.DoneList{},
		&models.Category{},
		&models.DoneListCategory{},
		&models.User{},
	)
}
