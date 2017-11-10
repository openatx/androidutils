package androidutils

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

func HWAddrWLAN() (string, error) {
	macaddr, err := getHWAddrWLAN()
	return strings.ToLower(macaddr), err
}

// WLAN Hardware Address, also known as wlan0 mac address
// Thanks to this article https://android.stackexchange.com/questions/142606/how-can-i-find-my-mac-address/142630#142630
func getHWAddrWLAN() (string, error) {
	// method 1
	if macaddr := GetProperty("ro.boot.wifimacaddr"); macaddr != "" {
		return macaddr, nil
	}
	// method 2
	data, err := ioutil.ReadFile("/sys/class/net/wlan0/address")
	if err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// method 3
	output, err := runShell("ip", "address", "show", "wlan0")
	if err != nil {
		return "", err
	}
	matches := regexp.MustCompile(`link/ether ([\w\d:]{17})`).FindStringSubmatch(output)
	if matches == nil {
		return "", errors.New("no mac address founded")
	}
	return matches[1], nil
}
