package controllers

import (
	"rogxd/rogxd/bus"
)

type ChargeController struct {
	Object *RogXObject
}

// Limit Get current charge limit.
func (controller *ChargeController) Limit(limit int16) {
	controller.Object.CallWithOut("Limit", &limit, -1)
	return
}

// SetLimit Set charge limit, the passed parameter must not below 0 and above 100, returns true if the operation was successful.
func (controller *ChargeController) SetLimit(limit byte) bool {
	return controller.Object.Call("SetLimit", limit) == nil
}

// WatchChargeLimit Register a handler, calling passed handler whenever the limit changes.
func (controller *ChargeController) WatchChargeLimit(handler func(limit byte)) bool {
	return controller.Object.RogX.RegisterHandler(bus.AsusDest, bus.DefaultInterface, bus.ChargeControllerPath, "NotifyCharge", func(body []interface{}) {
		handler(body[0].(byte))
	})
}

func GetChargeController(conn *RogXConn) *ChargeController {
	return &ChargeController{
		Object: NewObject(conn, bus.AsusDest, bus.DefaultInterface, bus.ChargeControllerPath),
	}
}
