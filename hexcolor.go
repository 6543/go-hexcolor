package hexcolor

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// HexColor represents a single hex color
type HexColor struct {
	original string
	hex      string
}

// NewHexColor convert string into a HexColor
func NewHexColor(hc string) (*HexColor, error) {
	c := &HexColor{original: hc}
	hc = strings.TrimLeft(strings.TrimSpace(strings.ToLower(hc)), "#")

	// normalize hex color
	if _, err := strconv.ParseUint(hc, 16, 24); err == nil {
		if len(hc) == 6 {
			c.hex = hc
			return c, nil
		}
		if len(hc) == 3 {
			c.hex = string([]byte{hc[0], hc[0], hc[1], hc[1], hc[2], hc[2]})
			return c, nil
		}
	}

	// resolve named color to hex color
	val, exist := CSS3ColorMap[hc]
	if exist {
		c.hex = string(val)
		return c, nil
	}

	return nil, fmt.Errorf("Malformed color: %s", hc)
}

// ToString return normalized hex code of color
func (c *HexColor) ToString() string {
	if c == nil {
		return ""
	}
	return c.hex
}

// ToRGBA return RGBA color
func (c *HexColor) ToRGBA() (*color.RGBA, error) {
	if c == nil {
		return nil, nil
	}
	var r, g, b uint8
	if _, err := fmt.Sscanf(c.hex, "%2x%2x%2x", &r, &g, &b); err != nil {
		return nil, err
	}
	return &color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 0xff,
	}, nil
}

// ToHexColor convert RGBA to HexColor
func ToHexColor(c *color.RGBA) *HexColor {
	toS := func(i uint8) string {
		h := fmt.Sprintf("%x", i)
		if len(h) == 1 {
			h = "0" + h
		}
		return h
	}

	return &HexColor{
		original: fmt.Sprintf("%v", c),
		hex:      toS(c.R) + toS(c.G) + toS(c.B),
	}
}
