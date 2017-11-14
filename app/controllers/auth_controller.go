package controllers

import (
	"github.com/k-kurikuri/gogo-done/app/auth"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	*revel.Controller
}

func (c AuthController) Authenticate() revel.Result {
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")

	con := db.Connection()
	var user models.User
	con.First(&user, "email = ?", email)
	if user.Email == "" {
		panic("non exist user")
	}

	hashPassword, err := auth.Crypt(password)
	if err != nil {
		panic("fail password crypt")
	}

	// TODO: hashとパスワードの判定がおかしい
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		panic("password wrong")
	}

	// this is sample json
	response := make(map[string]interface{})
	response["success"] = true
	response["user"] = user
	return c.RenderJSON(response)
}
