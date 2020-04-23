package geolocate

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

// ListNetworksWin produces a string slice of nearby wifi bssid's
func ListNetworksWin() []string {
	cmd := exec.Command("netsh", "wlan", "show", "networks", "mode=bssid")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("unable to run listNetworks_Win")
	}
	findbssid := regexp.MustCompile(`BSSID \d+\s+:\s([a-zA-Z0-9:]+)`)
	matches := findbssid.FindAllStringSubmatch(out.String(), -1)
	results := make([]string, len(matches))
	fmt.Println(len(matches))
	for i, match := range matches {
		results[i] = match[1]
	}
	return results
}
