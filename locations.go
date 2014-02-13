package main

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

type WebPageTestLocations struct {
	StatusCode int32
	StatusText string
	Data       map[string]Location
}

func GetLocations(wpturl string) (location WebPageTestLocations) {

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
	return
}
