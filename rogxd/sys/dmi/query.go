package dmi

import (
	"os"
	"path"
	"strings"
)

func Prop(prop DmiProp) string {
	buf, e := os.ReadFile(path.Join("/", "sys", "class", "dmi", "id", string(prop)))
	if e != nil {
		return ""
	}
	return strings.TrimSpace(string(buf))
}
