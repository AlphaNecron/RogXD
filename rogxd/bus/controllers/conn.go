package controllers

import (
	"github.com/godbus/dbus/v5"
	"log"
	"rogxd/rogxd/bus"
)

type (
	RogXConn struct {
		Handlers map[string]func(body []interface{})
		System   *dbus.Conn
	}
)

func (conn *RogXConn) handleSignals() {
	c := make(chan *dbus.Signal)
	conn.System.Signal(c)
	go func() {
		for r := range c {
			if handler := conn.Handlers[r.Name]; handler != nil {
				handler(r.Body)
			}
		}
	}()
}

func NewRogXConn() (conn *RogXConn) {
	sysConn, e := dbus.ConnectSystemBus()
	if e != nil {
		log.Fatalln(e)
	}
	conn = &RogXConn{Handlers: make(map[string]func([]interface{})), System: sysConn}
	conn.handleSignals()
	return
}

func (conn *RogXConn) Close() bool {
	return conn.System.Close() == nil
}

func (conn *RogXConn) RegisterHandler(path dbus.ObjectPath, member string, handler func(body []interface{})) bool {
	sConn := conn.System
	if sConn.AddMatchSignal(dbus.WithMatchObjectPath(path),
		dbus.WithMatchInterface(bus.ASUS_DEST),
		dbus.WithMatchSender(bus.ASUS_DEST),
		dbus.WithMatchMember(member)) != nil {
		return false
	}
	conn.Handlers[bus.Dest(bus.ASUS_DEST, member)] = handler
	return true
}

func (conn *RogXConn) ChargeController() ChargeController {
	return GetChargeController(conn)
}

func (conn *RogXConn) LedController() LedController {
	return GetLedController(conn)
}
