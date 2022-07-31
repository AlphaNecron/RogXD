package models

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

type LedMode byte

const (
	Static AuraMode = iota
	Breathe
	Strobe
	Rainbow
	Star
	Rain
	Highlight
	Laser
	Ripple
	Pulse
	Comet
	Flash
)

type AuraMode byte
