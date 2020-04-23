package main

import (
	"fmt"
	"micaiahwallace/geolocate"
)

func main() {
	nets := geolocate.ListNetworksWin()
	for _, net := range nets {
		fmt.Println("bssid: (", net, ")")
	}
}
