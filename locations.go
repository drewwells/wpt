package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	Browser       string
	Label         string
	PendingTests  map[string]int32
	Location      string
	RelayLocation string
	RelayServer   string
}

type PLocations struct {
	StatusCode int32
	StatusText string
	Data       map[string]Location
}

// Locations queries WebPageTest for information about the
// current available browsers.
func Locations(wpturl string) (location PLocations) {

	res, err := http.Get(wpturl + "/getLocations.php?f=json")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(robots, &location)

	if err != nil {
		log.Fatal(err)
	}
	return location
}
