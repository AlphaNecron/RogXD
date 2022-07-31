package ota

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"rogxd/rogxd/bus/controllers"
)

type (
	RogXServer struct {
		Iris *iris.Application
		RogX *controllers.RogXConn
	}
)

func (server RogXServer) Start() {
	_ = server.Iris.Listen(":7272")
}

func New(conn *controllers.RogXConn) (server *RogXServer) {
	app := iris.New()
	server = &RogXServer{
		Iris: app,
		RogX: conn,
	}
	for _, party := range Parties {
		server.RegisterParty(party)
	}
	return
}

func (server *RogXServer) RegisterParty(party Party) {
	app := server.Iris
	app.PartyFunc(party.Path, func(p router.Party) {
		for _, route := range party.Routes {
			route.Register(p, server.RogX)
		}
	})
}
