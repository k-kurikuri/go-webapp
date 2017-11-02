package db

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func Connection() *gorm.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbOption := "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", user+":"+pass+"@/"+dbName+dbOption)
	if err != nil {
		panic("Don't database open")
	}

	return db
}
