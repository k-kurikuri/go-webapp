package controllers

import (
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
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Create)
	}

	return c.RenderHTML("Register")
}

func (c User) validate(email, password, name string) {
	c.Validation.Email(email).Message("Please format email")

	c.Validation.MinSize(password, 8).Message("Password must be at least 8 characters")
	c.Validation.MaxSize(password, 16).Message("Password must be 16 characters or less")

	c.Validation.Required(name).Message("must be user name")
}
