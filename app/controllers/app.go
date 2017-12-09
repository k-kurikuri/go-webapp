package controllers

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/filters"
	"github.com/k-kurikuri/gogo-done/app/models"
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
	res["title"] = "映画見た"
	res["category"] = "Movie"

	return c.RenderJSON(res)
}

func (c App) sessionUser() (models.User, error) {
	jsonStr := c.Session["user"]

	jsonBytes := ([]byte)(jsonStr)

	user := models.User{}

	err := json.Unmarshal(jsonBytes, &user)

	return user, err
}
