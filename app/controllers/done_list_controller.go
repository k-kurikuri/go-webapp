package controllers

import (
	"github.com/revel/revel"
)

type DoneListController struct {
	*revel.Controller
}

func (c DoneListController) Index() revel.Result {
	return c.Render()
}
