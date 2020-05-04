package geolocate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

// ListNetworksWin produces a string slice of nearby wifi bssid's
func ListNetworksWin() []MacAddr {

	// execute netsh command
	cmd := exec.Command("netsh", "wlan", "show", "networks", "mode=bssid")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("unable to run listNetworks_Win")
	}

	// use regex to find bssid lines
	findbssid := regexp.MustCompile(`BSSID \d+\s+:\s([a-zA-Z0-9:]+)`)
	matches := findbssid.FindAllStringSubmatch(out.String(), -1)

	// map the bssids into a macaddr slice
	results := make([]MacAddr, len(matches))
	for i, match := range matches {
		results[i] = MacAddr{Mac: match[1]}
	}
	return results
}

// Locate requests current device location based on nearby wifi BSSID's
func Locate() (string, error) {

	// Get list of wifi bssid's
	wlanlist := ListNetworksWin()

	// ensure at least 1 bssid exists
	if len(wlanlist) < 1 {
		return "", errors.New("Less than 1 bssid available")
	}

	// create json request payload
	req := LocateRequest{Wlan: wlanlist}
	reqjson, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// set fallback for single wifi case
	fallbackStr := ""
	if len(wlanlist) == 1 {
		fallbackStr = "&fallback=singleWifi"
	}

	// submit post request
	resp, err := http.Post("https://pos.ls.hereapi.com/positioning/v1/locate?apiKey=***REMOVED***"+fallbackStr, "application/json", bytes.NewBuffer(reqjson))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// create struct from response
	locateres := LocateResp{}
	err = json.Unmarshal(body, &locateres)
	if err != nil {
		return "", err
	}

	// check for api errors
	if locateres.Error.Code != 0 {
		return "", errors.New(locateres.Error.Description)
	}

	// json marshal location object
	locjson, err := json.Marshal(locateres.Location)
	if err != nil {
		return "", err
	}

	// return location
	return string(locjson), nil
}
