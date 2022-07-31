package models

import (
	"encoding/json"
	"fmt"
)

type Color struct {
	R byte // Red channel
	G byte // Green channel
	B byte // Blue channel
}

func NewColor(colorArr interface{}) Color {
	c := colorArr.([]interface{})
	r, g, b := c[0].(byte), c[1].(byte), c[2].(byte)
	return Color{r, g, b}
}

func (c Color) Hex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Hex())
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
