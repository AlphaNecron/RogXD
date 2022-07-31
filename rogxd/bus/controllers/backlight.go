package controllers

import "rogxd/rogxd/bus"

type BacklightController struct {
	Object *RogXObject
}

func GetBacklightController(conn *RogXConn) *BacklightController {
	return &BacklightController{
		Object: NewObject(conn, bus.AsusDest, bus.DefaultInterface, bus.ChargeControllerPath),
	}
}
