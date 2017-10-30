package controllers

import (
	"github.com/revel/revel"
)

type Corgi struct {
	*revel.Controller
}

func (c Corgi) Index() revel.Result {
	dogName := "this is corgi"
	return c.Render(dogName)
}
