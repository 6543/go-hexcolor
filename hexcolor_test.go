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
		{"15di", "1155dd", true},
		{"2e8b57", "2e8b57", false},
		{"2e8b572e8b57", "", true},

		{"Red", "ff0000", false},
		{"blue", "0000ff", false},
		{"reed", "", true},
		{"", "", true},
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
