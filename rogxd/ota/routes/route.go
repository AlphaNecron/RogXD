package routes

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"rogxd/rogxd/bus/controllers"
)

const (
	Get   Method = "GET"
	Post  Method = "POST"
	Patch Method = "PATCH"
)

type (
	Method string
	Route  struct {
		Method  Method
		Path    string
		Handler func(ctx Context)
	}
)

func (route Route) Register(party router.Party, conn *controllers.RogXConn) *router.Route {
	return party.Handle(string(route.Method), route.Path, func(baseCtx context.Context) {
		route.Handler(Context{
			Base: &baseCtx,
			RogX: conn,
		})
	})
}
