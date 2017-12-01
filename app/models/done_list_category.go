package models

type DoneListCategory struct {
	Id         uint `gorm:"primary_key"`
	DoneListId uint
	CategoryId uint
}
