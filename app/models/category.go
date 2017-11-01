package models

type Category struct {
	Id        uint `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
}
