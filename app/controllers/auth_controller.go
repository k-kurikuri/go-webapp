package controllers

import (
	"github.com/revel/revel"
)

type AuthController struct {
	*revel.Controller
}

func (c AuthController) Authenticate() revel.Result {
	// this is sample json
	response := make(map[string]interface{})
	response["user"] = "authUser"

	return c.RenderJSON(response)
}
