package filters

import (
	"github.com/revel/revel"
)

var AuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	var user string = c.Session["user"]
	if user == "" {
		panic("not authenticated")
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
