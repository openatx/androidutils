package androidutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotationDumpsysInput(t *testing.T) {
	r, err := getRotationDumpsysInput(`SurfaceWidth: 1080px
		SurfaceHeight: 1920px
		SurfaceLeft: 0
		SurfaceTop: 0
		SurfaceOrientation: 1
		Translation and Scaling Factors:`)
	assert.NoError(t, err)
	assert.Equal(t, 90, r)
}

func TestRotationMinicap(t *testing.T) {
	r, err := getRotationMinicap(`{
		"id": 0,
		"width": 1080,
		"height": 1920,
		"xdpi": 442.45,
		"ydpi": 443.35,
		"size": 4.97,
		"density": 3.00,
		"fps": 60.00,
		"secure": true,
		"rotation": 90
	}`)
	assert.NoError(t, err)
	assert.Equal(t, 90, r)
}
