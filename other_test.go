package androidutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWmSize(t *testing.T) {
	output := `Physical size: 1440x2560
Override size: 1080x1920`
	w, h, err := parseWmSize(output)
	assert.NoError(t, err)
	assert.Equal(t, w, 1080)
	assert.Equal(t, h, 1920)

	output = `Physical size: 1440x2560`
	w, h, err = parseWmSize(output)
	assert.NoError(t, err)
	assert.Equal(t, w, 1440)
	assert.Equal(t, h, 2560)

	output = `Exit 0`
	_, _, err = parseWmSize(output)
	assert.Error(t, err)
}
