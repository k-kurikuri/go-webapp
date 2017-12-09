package controllers

import (
	"github.com/k-kurikuri/gogo-done/app/filters"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func init() {
	revel.FilterController(App{}).Insert(filters.AuthFilter, revel.BEFORE, revel.ActionInvoker)
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Create() revel.Result {
	// TODO:
	res := make(map[string]interface{})
	res["title"] = c.Params.Form.Get("title")
	res["category"] = c.Params.Form.Get("category")

	return c.RenderJSON(res)
}
