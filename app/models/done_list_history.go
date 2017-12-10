package models

import "time"

type DoneListHistory struct {
	Id         uint   `gorm:"primary_key"`
	DoneListId uint   `gorm:"index"`
	Note       string `gorm:"size:255"`
	PostedAt   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
