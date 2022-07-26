package bus

import (
	"encoding/json"
	"github.com/godbus/dbus/v5"
)

func CallWithOut(obj dbus.BusObject, method string, out interface{}, fallback interface{}, args ...interface{}) {
	call := obj.Call(ASUS_DEST+"."+method, 0, args...)
	if call.Store(out) != nil {
		out = fallback
	}
}

func Call(obj dbus.BusObject, method string, args ...interface{}) error {
	call := obj.Call(ASUS_DEST+"."+method, 0, args...)
	return call.Err
}

func Dest(base string, prop string) string {
	return base + "." + prop
}

func Path(path dbus.ObjectPath, member string) string {
	return string(path) + "." + member
}

func Deserialize(inp string, out any) bool {
	if json.Unmarshal([]byte(inp), out) != nil {
		return false
	}
	return true
}

func Property(obj dbus.BusObject, prop string, out interface{}) bool {
	if obj.StoreProperty(Dest(ASUS_DEST, prop), out) == nil {
		return true
	}
	return false
}
