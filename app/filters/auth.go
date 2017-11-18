package filters

import (
	"fmt"
	"github.com/revel/revel"
)

var AuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	// TODO: Filtering
	fmt.Print("Hello AuthFilter!")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
