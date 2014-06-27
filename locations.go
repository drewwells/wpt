package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kr/pretty"
)

type jloc struct {
	Browser       string
	Label         string `json:"label"`
	PendingTests  map[string]int32
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
	Total   int32
	Testing int32
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

	var jlocs jlocs

	err := json.Unmarshal(bs, &jlocs)

	if err != nil {
		log.Fatal(err)
	}

	locs := make([]Location, len(jlocs.Data))
	i := 0
	for _, v := range jlocs.Data {
		locs[i].Name = v.Location
		locs[i].Browser = v.Browser
		locs[i].Label = v.Label
		locs[i].Total = v.PendingTests["Total"]
		locs[i].Testing = v.PendingTests["Testing"]
		locs[i].Busy = v.PendingTests["Idle"] != 1
		i = i + 1
	}
	pretty.Println(locs)
	return locs, err

}
