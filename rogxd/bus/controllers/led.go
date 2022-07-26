package controllers

import (
	"encoding/json"
	"rogxd/rogxd/bus"
)

var LedModes = [12]string{
	"Static",
	"Breathe",
	"Strobe",
	"Rainbow",
	"Star",
	"Rain",
	"Highlight",
	"Laser",
	"Ripple",
	"Pulse",
	"Comet",
	"Flash",
}

const (
	Off Brightness = iota
	Low
	Medium
	High
)

const (
	Left  Direction = "Left"
	Right           = "Right"
	Top             = "Top"
	Down            = "Down"
)

type (
	Direction     string
	Brightness    uint32
	LedController struct {
		NextLedBrightness func() bool                               // Cycle to the next brightness level.
		PrevLedBrightness func() bool                               // Cycle to the previous brightness level.
		NextLedMode       func() bool                               // Cycle to the next led mode.
		PrevLedMode       func() bool                               // Cycle to the previous led mode.
		LedBrightness     func() int16                              // Get the current keyboard brightness.
		LedMode           func() LedMode                            // Get the current led mode.
		LedModes          func() map[string]LedMode                 // List all available led modes.
		SetLedBrightness  func(brightness Brightness) bool          // Set a specific keyboard brightness.
		WatchLedChange    func(handler func(signal LedSignal)) bool // Subscribe for led mode changes.
	}
	LedMode struct {
		Mode      string `json:"mode"`
		Zone      string `json:"zone"`
		Color1    Color  `json:"colour1"`
		Color2    Color  `json:"colour2"`
		Speed     string `json:"speed"`
		Direction string `json:"direction"`
	}
	Color struct {
		R byte // Red channel
		G byte // Green channel
		B byte // Blue channel
	}
	LedSignal struct {
		Mode      uint32
		Zone      uint32
		Color1    Color
		Color2    Color
		Speed     uint32
		Direction uint32
	}
)

func ParseLedSignal(signal []interface{}) LedSignal {
	if len(signal) < 5 {
		return LedSignal{}
	}
	return LedSignal{signal[0].(uint32), signal[1].(uint32), NewColor(signal[2]), NewColor(signal[3]), signal[4].(uint32), signal[5].(uint32)}
}

func NewColor(colorArr interface{}) Color {
	c := colorArr.([]interface{})
	r, g, b := c[0].(byte), c[1].(byte), c[2].(byte)
	return Color{r, g, b}
}

func (c *Color) UnmarshalJSON(bytes []byte) error {
	var cb [3]byte
	err := json.Unmarshal(bytes, &cb)
	if err != nil {
		return err
	}
	c.R, c.G, c.B = cb[0], cb[1], cb[2]
	return nil
}

func GetLedController(conn *RogXConn) LedController {
	obj := conn.System.Object(bus.ASUS_DEST, bus.LED_CONTROLLER_PATH)
	return LedController{
		NextLedBrightness: func() bool {
			return bus.Call(obj, "NextLedBrightness") == nil
		},
		PrevLedBrightness: func() bool {
			return bus.Call(obj, "PrevLedBrightness") == nil
		},
		NextLedMode: func() bool {
			return bus.Call(obj, "NextLedMode") == nil
		},
		PrevLedMode: func() bool {
			return bus.Call(obj, "PrevLedMode") == nil
		},
		LedBrightness: func() (brightness int16) {
			bus.Property(obj, "LedBrightness", &brightness)
			return
		},
		LedMode: func() (mode LedMode) {
			var modeStr string
			if !bus.Property(obj, "LedMode", &modeStr) {
				return
			}
			bus.Deserialize(modeStr, &mode)
			return
		},
		LedModes: func() (modes map[string]LedMode) {
			var modesStr string
			if !bus.Property(obj, "LedModes", &modesStr) {
				return
			}
			bus.Deserialize(modesStr, &modes)
			return
		},
		SetLedBrightness: func(brightness Brightness) bool {
			return bus.Call(obj, "SetBrightness", brightness) == nil
		},
		WatchLedChange: func(handler func(signal LedSignal)) bool {
			return conn.RegisterHandler(bus.LED_CONTROLLER_PATH, "NotifyLed", func(body []interface{}) {
				signal := ParseLedSignal(body[0].([]interface{}))
				handler(signal)
			})
		},
	}
}
