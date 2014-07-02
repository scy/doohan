package app

import "github.com/revel/revel"

func init() {
	revel.Filters = []revel.Filter{
		revel.PanicFilter,
		revel.RouterFilter,
		revel.ParamsFilter,
		revel.ActionInvoker,
	}
}
