package controllers

import (
	"github.com/k-kurikuri/gogo-done/app/auth"
	"github.com/k-kurikuri/gogo-done/app/db"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
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
		// TODO: validation logic
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Create)
	}

	// email登録チェック
	user := models.User{}
	con := db.Connection()
	con.Where("email = ?", email).FirstOrInit(&user)

	if user.Email == email {
		c.Flash.Error("This is a registered email address")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Create)
	}

	hashPass, err := auth.Crypt(password)
	if err != nil {
		c.Flash.Error("Create Hash Password Failed")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Create)
	}

	user = models.User{Email: email, HashPass: hashPass, Name: name}
	con.Create(&user)

	return c.RenderHTML("Register")
}

func (c User) validate(email, password, name string) {
	c.Validation.Email(email).Message("Please format email")

	c.Validation.MinSize(password, 8).Message("Password must be at least 8 characters")
	c.Validation.MaxSize(password, 16).Message("Password must be 16 characters or less")

	c.Validation.Required(name).Message("must be user name")
}
