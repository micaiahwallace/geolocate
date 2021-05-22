# geolocate
## Summary

This utility provides a quick way to get a JSON string containing latitude, longitude and accuracy for a windows device with a WiFi card.

This utility is only for Windows at the moment and does not have any external go dependencies, it just requires the `netsh wlan` command and access to hereapi with an API token.

## Usage

Get running quick by installing with go get to have access to the geolocate command
```
C:\> go get github.com/micaiahwallace/geolocate/cmd/geolocate
C:\> geolocate.exe --help
```

Use the `--help` flag to print the command usage
```shell
C:\>geolocate.exe --help
Locate current device using nearby WAPs
  -apikey string
        Here location service api key
  -out string
        Output file to write json location
```

## Example

Running the command with the API key specified returns the JSON data directly to stdout
```shell
C:\>geolocate.exe -apikey REPLACE_WITH_HEREAPI_KEY
{"lat":30.123,"lng":-103.123,"accuracy":641}
```

You can store the results in a file with either the `-out` option or shell redirection. Shell redirection is useful if you want to setup a windows Scheduled Task and log the location on a schedule by appending to an existing log file.
```shell
C:\>geolocate.exe -apikey REPLACE_WITH_HEREAPI_KEY -out current_location.json
C:\>geolocate.exe -apikey >> location_log.json
```

