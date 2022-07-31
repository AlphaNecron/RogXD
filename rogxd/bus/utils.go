package bus

import (
	"encoding/json"
	"github.com/godbus/dbus/v5"
)

func Dest(base string, prop string) string {
	return base + "." + prop
}

func Path(path dbus.ObjectPath, member string) string {
	return string(path) + "." + member
}

func Deserialize(inp string, out any) bool {
	if e := json.Unmarshal([]byte(inp), out); e != nil {
		return false
	}
	return true
}
