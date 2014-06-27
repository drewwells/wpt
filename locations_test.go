package wpt

import (
	"io/ioutil"
	"testing"
)

func TestLocations(t *testing.T) {

	//Reading Test Data

	bytes, err := ioutil.ReadFile("./test/location.json")

	if err != nil {
		t.Errorf(err.Error())
	}

	locs, err := processLoc(bytes)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%v", locs)
}
