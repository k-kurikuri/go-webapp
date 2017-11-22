package controllers

import (
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Create() revel.Result {
	return c.RenderHTML("Create")
}

func (c User) Register() revel.Result {
	return c.RenderHTML("Register")
}
