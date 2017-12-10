package models

import (
	"time"
)

type DoneList struct {
	Id         uint   `gorm:"primary_key"`
	UserId     uint   `gorm:"index"`
	Title      string `gorm:"size:255"`
	CategoryId uint   `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
