package controllers

import (
	"rogxd/rogxd/bus"
)

type ChargeController struct {
	Limit            func() int16                        // Get current charge limit.
	SetLimit         func(limit byte) bool               // Set charge limit, the passed parameter must not below 0 and above 100, returns true if the operation was successful.
	WatchChargeLimit func(handler func(limit byte)) bool // Register a handler, calling passed handler whenever the limit changes.
}

func GetChargeController(conn *RogXConn) ChargeController {
	obj := conn.System.Object(bus.ASUS_DEST, bus.CHARGE_CONTROLLER_PATH)
	return ChargeController{
		Limit: func() (limit int16) {
			bus.CallWithOut(obj, "Limit", &limit, -1)
			return
		},
		SetLimit: func(limit byte) bool {
			return bus.Call(obj, "SetLimit", limit) == nil
		},
		WatchChargeLimit: func(handler func(limit byte)) bool {
			return conn.RegisterHandler(bus.CHARGE_CONTROLLER_PATH, "NotifyCharge", func(body []interface{}) {
				handler(body[0].(byte))
			})
		},
	}
}
