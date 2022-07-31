package controllers

import (
	"rogxd/rogxd/bus"
	"rogxd/rogxd/bus/models"
)

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
	Brightness    byte
	LedController struct {
		Object          *RogXObject
		BacklightObject *RogXObject
	}
	LedMode struct {
		Mode      string       `json:"mode"`
		Zone      string       `json:"zone"`
		Color1    models.Color `json:"colour1"`
		Color2    models.Color `json:"colour2"`
		Speed     string       `json:"speed"`
		Direction string       `json:"direction"`
	}
	AuraEffect struct {
		Mode      models.AuraMode
		Zone      uint32
		Color1    models.Color
		Color2    models.Color
		Speed     uint32
		Direction uint32
	}
)

func ParseAuraEffect(effect []interface{}) AuraEffect {
	if len(effect) < 5 {
		return AuraEffect{}
	}
	return AuraEffect{models.AuraMode(effect[0].(uint32)), effect[1].(uint32), models.NewColor(effect[2]), models.NewColor(effect[3]), effect[4].(uint32), effect[5].(uint32)}
}

// NextLedBrightness Cycle to the next brightness level.
func (controller *LedController) NextLedBrightness() bool {
	return controller.Object.Call("NextLedBrightness") == nil
}

// PrevLedBrightness Cycle to the previous brightness level.
func (controller *LedController) PrevLedBrightness() bool {
	return controller.Object.Call("PrevLedBrightness") == nil
}

// NextLedMode Cycle to the next led mode.
func (controller *LedController) NextLedMode() bool {
	return controller.Object.Call("NextLedMode") == nil
}

// PrevLedMode Cycle to the previous led mode.
func (controller *LedController) PrevLedMode() bool {
	return controller.Object.Call("PrevLedMode") == nil
}

// LedBrightness Get the current keyboard brightness.
func (controller *LedController) LedBrightness() Brightness {
	var brightness int32
	controller.BacklightObject.CallWithOut("GetBrightness", &brightness, 0)
	return Brightness(brightness)
}

// LedMode Get the current led mode.
func (controller *LedController) LedMode() (mode LedMode) {
	var modeStr string
	if !controller.Object.Property("LedMode", &modeStr) {
		return
	}
	bus.Deserialize(modeStr, &mode)
	return
}

// LedModes List all available led modes.
func (controller *LedController) LedModes() (modes map[string]LedMode) {
	var modesStr string
	if !controller.Object.Property("LedModes", &modesStr) {
		return
	}
	bus.Deserialize(modesStr, &modes)
	return
}

// SetLedBrightness Set a specific keyboard brightness.
func (controller *LedController) SetLedBrightness(brightness Brightness) bool {
	return controller.BacklightObject.Call("SetBrightness", int32(brightness)) == nil
}

// WatchLedChange Subscribe for led mode changes.
func (controller *LedController) WatchLedChange(handler func(effect AuraEffect)) bool {
	return controller.Object.RogX.RegisterHandler(bus.AsusDest, bus.DefaultInterface, bus.LedControllerPath, "NotifyLed", func(body []interface{}) {
		auraEffect := ParseAuraEffect(body[0].([]interface{}))
		handler(auraEffect)
	})
}

// MaxBrightness Get maximum brightness
func (controller *LedController) MaxBrightness() (brightness Brightness) {
	controller.BacklightObject.CallWithOut("GetMaxBrightness", &brightness, 0)
	return
}

// WatchBrightnessChange Subscribe for brightness changes
func (controller *LedController) WatchBrightnessChange(handler func(brightness Brightness)) {
	controller.Object.RogX.RegisterHandler(bus.PowerDest, bus.KbdBacklightInterface, bus.KbdBacklightControllerPath, "BrightnessChanged", func(body []interface{}) {
		handler(Brightness(body[0].(int32)))
	})
}

func GetLedController(conn *RogXConn) *LedController {
	return &LedController{
		Object:          NewObject(conn, bus.AsusDest, bus.DefaultInterface, bus.LedControllerPath),
		BacklightObject: NewObject(conn, bus.PowerDest, bus.KbdBacklightInterface, bus.KbdBacklightControllerPath),
	}
}
