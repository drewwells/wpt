package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var wpturl = "http://webpagetest.eng.wsm.local/getLocations.php?f=json"

type Location struct {
	Browser       string
	Label         string
	PendingTests  map[string]int32
	Location      string
	RelayLocation string
	RelayServer   string
}

func main() {
	type WebPageTestLocations struct {
		StatusCode int32
		StatusText string
		Data       map[string]Location
	}
	var location WebPageTestLocations
	res, err := http.Get(wpturl)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(robots, &location)
	fmt.Printf("%+v", location)
}
