package controllers

import (
	"github.com/godbus/dbus/v5"
	"rogxd/rogxd/bus"
)

type RogXObject struct {
	Dest      string
	Path      dbus.ObjectPath
	Interface string
	Base      dbus.BusObject
	RogX      *RogXConn
}

func (obj *RogXObject) CallWithOut(method string, out interface{}, fallback interface{}, args ...interface{}) {
	call := obj.Base.Call(bus.Dest(obj.Interface, method), 0, args...)
	if e := call.Store(out); e != nil {
		out = fallback
	}
}

func NewObject(conn *RogXConn, dest string, iface string, path dbus.ObjectPath) *RogXObject {
	return &RogXObject{
		Dest:      dest,
		Path:      path,
		Interface: iface,
		Base:      conn.Bus.Object(dest, path),
		RogX:      conn,
	}
}

func (obj *RogXObject) Call(method string, args ...interface{}) error {
	call := obj.Base.Call(bus.Dest(obj.Interface, method), 0, args...)
	return call.Err
}

func (obj *RogXObject) Property(prop string, out interface{}) bool {
	return obj.Base.StoreProperty(bus.Dest(obj.Interface, prop), out) == nil
}
