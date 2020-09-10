package go_hexcolor

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
)

// HexColor represents a single hex color
type HexColor struct {
	original string
	hex      string
}

// The compiled regular expression used to test the validity of a color statement
var (
	hexDefault *regexp.Regexp
	hexShort   *regexp.Regexp
)

// The raw regular expression string used for testing the validity
const (
	hexDefaultRaw string = `[0-9a-f]{6}`
	hexShortRaw   string = `[0-9a-f]{3}`
)

func init() {
	hexDefault = regexp.MustCompile("^" + hexDefaultRaw + "$")
	hexShort = regexp.MustCompile("^" + hexShortRaw + "$")
}

// NewHexColor convert string into a HexColor
func NewHexColor(hc string) (*HexColor, error) {
	c := &HexColor{original: hc}
	hc = strings.TrimLeft(strings.ToLower(hc), "#")

	if hexDefault.MatchString(hc) {
		c.hex = hc
		return c, nil
	}
	if hexShort.MatchString(hc) {
		c.hex = hc[:1] + hc[:1] + hc[1:][:1] + hc[1:][:1] + hc[2:] + hc[2:]
		return c, nil
	}

	// handle named colors

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
	rr, err := strconv.ParseUint(c.hex[:2], 16, 8)
	if err != nil {
		return nil, err
	}
	gg, err := strconv.ParseUint(c.hex[2:][:2], 16, 8)
	if err != nil {
		return nil, err
	}
	bb, err := strconv.ParseUint(c.hex[4:], 16, 8)
	if err != nil {
		return nil, err
	}
	return &color.RGBA{
		R: uint8(rr),
		G: uint8(gg),
		B: uint8(bb),
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
