package ui

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
)

func TransformColorRGBA(name string) (color.RGBA, error) {
	if c, ok := colornames.Map[name]; ok {
		return c, nil
	}

	return color.RGBA{0x00, 0x00, 0x00, 0xff}, fmt.Errorf("color %s not found", name)
}

func TransformColorHex(name string) (string, error) {
	if c, ok := colornames.Map[name]; ok {
		return fmt.Sprintf("#%02X%02X%02X%02X", c.R, c.G, c.B, c.A), nil
	}

	return "", fmt.Errorf("color %s not found", name)
}
