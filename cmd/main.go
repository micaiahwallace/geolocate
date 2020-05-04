package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"micaiahwallace/geolocate"
	"os"
)

// Usage prints cli usage
func usage() {
	fmt.Println("Locate current device using nearby WAPs")
	flag.PrintDefaults()
}

func main() {

	// get cli args
	var outfile string

	// parse cli args
	flag.StringVar(&outfile, "out", "", "Output file to write json location")
	flag.CommandLine.Usage = usage
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	// Fetch the location
	location, err := geolocate.Locate()
	if err != nil {
		log.Println("Failed to locate", err.Error())
	}

	// If file specified, write to file instead of stdout
	if len(outfile) > 0 {
		writeFile(location, outfile)
	} else {
		fmt.Println(location)
	}
}

func writeFile(txt, file string) error {
	return ioutil.WriteFile(file, []byte(txt), 0755)
}
