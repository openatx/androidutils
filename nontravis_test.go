// +build windows

package androidutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunShell(t *testing.T) {
	output, err := runShell("echo", "-n", "hello")
	assert.NoError(t, err)
	assert.Equal(t, "hello", output)
}
