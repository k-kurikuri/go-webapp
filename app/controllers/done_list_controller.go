package controllers

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/filters"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
)

type DoneListController struct {
	*revel.Controller
}

func init() {
	revel.FilterController(DoneListController{}).Insert(filters.AuthFilter, revel.BEFORE, revel.ActionInvoker)
}

func (c DoneListController) Index() revel.Result {
	user, err := c.sessionUser()
	if err != nil {
		panic("json parse error")
	}

	return c.Render(user)
}

func (c DoneListController) sessionUser() (*models.User, error) {
	jsonStr := c.Session["user"]

	jsonBytes := ([]byte)(jsonStr)

	user := new(models.User)

	err := json.Unmarshal(jsonBytes, user)

	return user, err
}
