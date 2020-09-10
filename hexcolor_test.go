package hexcolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHexColor(t *testing.T) {
	cases := []struct {
		color string
		hex   string
		err   bool
	}{
		{"#dd22cc", "dd22cc", false},
		{"#adf", "aaddff", false},
		{"15d", "1155dd", false},
		{"a5i", "", true},
	}

	for _, tc := range cases {
		hc, err := NewHexColor(tc.color)
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			if assert.NotNil(t, hc) {
				assert.EqualValues(t, tc.hex, hc.hex)
			}
		}
	}
}
