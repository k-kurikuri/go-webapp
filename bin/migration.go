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
		&models.DoneListHistory{},
		&models.DoneList{},
		&models.Category{},
		&models.User{},
	)

	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.DoneList{},
		&models.Category{},
		&models.DoneListHistory{},
		&models.User{},
	)

	// params : foreignkey field, destination table(id) ,ONDELETE ,ONUPDATE
	db.Model(&models.DoneList{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.DoneListHistory{}).AddForeignKey("done_list_id", "done_lists(id)", "RESTRICT", "RESTRICT")
}
