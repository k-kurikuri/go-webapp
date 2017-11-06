package models

import (
	"github.com/revel/modules/auth/driver/secret"
	_ "golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	HashPass  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	secret.BcryptAuth `sql:"-"`
}
