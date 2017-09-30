package androidutils

import (
	"io/ioutil"
	"strings"
)

// WLAN Hardware Address, also known as wlan0 mac address
func HWAddrWLAN() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/net/wlan0/address")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
