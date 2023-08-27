package gee

import (
	"gee/router"
)

type Context = router.Context

func Default() *router.Engine {
	return router.Default()
}