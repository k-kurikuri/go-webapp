package controllers

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/auth"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
	"log"
)

type User struct {
	*revel.Controller
}

func (c User) Create() revel.Result {
	return c.Render()
}

func (c User) Register() revel.Result {
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")
	name := c.Params.Form.Get("name")

	c.validate(email, password, name)

	if c.Validation.HasErrors() {
		return c.flashWithRedirect("")
	}

	// email登録チェック
	user := models.User{}
	con := db.Connection()
	con.Where("email = ?", email).FirstOrInit(&user)

	if user.Email == email {
		return c.flashWithRedirect("This is a registered email address")
	}

	hashPass, err := auth.Crypt(password)
	if err != nil {
		return c.flashWithRedirect("Create Hash Password Failed")
	}

	user = models.User{Email: email, HashPass: hashPass, Name: name}
	con.Create(&user)

	c.authSession(user)

	return c.Redirect(App.Index)
}

func (c User) validate(email, password, name string) {
	c.Validation.Email(email).Message("Please format email")

	c.Validation.MinSize(password, 8).Message("Password must be at least 8 characters")
	c.Validation.MaxSize(password, 16).Message("Password must be 16 characters or less")

	c.Validation.Required(name).Message("must be user name")
}

func (c User) flashWithRedirect(errMsg string) revel.Result {
	c.Flash.Error(errMsg)
	c.Validation.Keep()
	c.FlashParams()
	return c.Redirect(User.Create)
}

func (c User) authSession(user models.User) {
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Panic("json marshal error")
	}
	c.Session["user"] = string(jsonBytes)
}
