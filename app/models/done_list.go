package models

import (
	"time"
)

type DoneList struct {
	Id        uint   `gorm:"primary_key"`
	UserId    uint   `gorm:"index"`
	Title     string `gorm:"size:255"`
	Note      string `gorm:"size:255"`
	PostedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
