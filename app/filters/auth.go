package filters

import (
	"encoding/json"
	"github.com/k-kurikuri/gogo-done/app/models"
	"github.com/revel/revel"
)

var AuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	var userStr string = c.Session["user"]
	if userStr == "" {
		panic("not authenticated")
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.

	user := models.User{}

	json.Unmarshal([]byte(userStr), &user)

	c.Render(user)
}
