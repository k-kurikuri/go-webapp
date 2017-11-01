package models

import "time"

type User struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
