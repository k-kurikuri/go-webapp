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
	user, err := c.sessionUser()
	if err != nil {
		panic("json parse error")
	}

	return c.Render(user)
}

func (c App) sessionUser() (*models.User, error) {
	jsonStr := c.Session["user"]

	jsonBytes := ([]byte)(jsonStr)

	user := new(models.User)

	err := json.Unmarshal(jsonBytes, user)

	return user, err
}
