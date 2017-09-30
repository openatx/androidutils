package androidutils

import (
	"os/exec"
	"regexp"
	"strings"
)

func runShell(args ...string) (out string, err error) {
	output, err := exec.Command("sh", "-c", strings.Join(args, " ")).CombinedOutput()
	return string(output), err
}

// Properties extract info from $ adb shell getprop
func Properties() (props map[string]string, err error) {
	propOutput, err := runShell("getprop")
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`\[(.*?)\]:\s*\[(.*?)\]`)
	matches := re.FindAllStringSubmatch(propOutput, -1)
	props = make(map[string]string)
	for _, m := range matches {
		var key = m[1]
		var val = m[2]
		props[key] = val
	}
	return
}
