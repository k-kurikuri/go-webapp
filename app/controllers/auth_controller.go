package controllers

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/k-kurikuri/gogo-done/app/routes"
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

	hashPassword := user.HashPass

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		panic("password wrong")
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic("json marshal error")
	}

	c.Session["user"] = string(jsonBytes)
	c.Flash.Success("Welcome " + user.Name)

	return c.Redirect(routes.DoneListController.Index())
}
