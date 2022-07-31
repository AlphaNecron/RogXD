package controllers

import (
	"github.com/godbus/dbus/v5"
	"log"
	"rogxd/rogxd/bus"
)

type (
	RogXConn struct {
		Handlers            map[string]func(body []interface{})
		Bus                 *dbus.Conn
		BacklightController *BacklightController
		LedController       *LedController
		ChargeController    *ChargeController
	}
)

func handleSignals(conn *RogXConn) {
	c := make(chan *dbus.Signal)
	conn.Bus.Signal(c)
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
	conn = &RogXConn{Handlers: make(map[string]func([]interface{})), Bus: sysConn}
	conn.LedController = GetLedController(conn)
	conn.ChargeController = GetChargeController(conn)
	conn.BacklightController = GetBacklightController(conn)
	handleSignals(conn)
	return
}

func (conn *RogXConn) Close() bool {
	return conn.Bus.Close() == nil
}

func (conn *RogXConn) RegisterHandler(dest string, iface string, path dbus.ObjectPath, member string, handler func(body []interface{})) bool {
	sConn := conn.Bus
	if sConn.AddMatchSignal(dbus.WithMatchObjectPath(path),
		dbus.WithMatchInterface(iface),
		dbus.WithMatchSender(dest),
		dbus.WithMatchMember(member)) != nil {
		return false
	}
	conn.Handlers[bus.Dest(iface, member)] = handler
	return true
}
