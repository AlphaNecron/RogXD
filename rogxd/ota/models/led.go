package models

const (
	Next Operation = "next"
	Prev           = "prev"
)

type (
	Operation          string
	BrightnessResponse struct {
		Brightness byte `json:"brightness"`
	}
	BrightnessRequest struct {
		Operation  Operation `json:"operation"`
		Brightness byte      `json:"brightness"`
	}
	LedMode struct {
		Mode string `json:"mode"`
	}
)
