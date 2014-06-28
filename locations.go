package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type jloc struct {
	Browser       string
	Label         string `json:"label"`
	PendingTests  map[string]int
	Location      string
	RelayLocation string
	RelayServer   string
}

type jlocs struct {
	StatusCode int32
	StatusText string
	Data       map[string]jloc
}

type Location struct {
	Name    string
	Browser string
	Label   string

	Busy    bool
	Total   int
	Testing int
}

// Locations queries WebPageTest for information about the
// current available browsers.
func Locations(url string) ([]Location, error) {

	res, err := http.Get(url + "/getLocations.php?f=json")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	return processLoc(body)
}

func processLoc(bs []byte) ([]Location, error) {

	var jls jlocs

	err := json.Unmarshal(bs, &jls)

	if err != nil {
		log.Fatal(err)
	}

	keys := make(sort.StringSlice, len(jls.Data))
	i := 0
	for j, _ := range jls.Data {
		keys[i] = j
		i = i + 1
	}

	keys.Sort()
	locs := make([]Location, len(jls.Data))
	for i, v := range keys {
		vv := jls.Data[v]
		locs[i].Name = vv.Location
		locs[i].Browser = vv.Browser
		locs[i].Label = vv.Label
		locs[i].Total = vv.PendingTests["Total"]
		locs[i].Testing = vv.PendingTests["Testing"]
		locs[i].Busy = vv.PendingTests["Idle"] != 0
	}

	return locs, err
}
