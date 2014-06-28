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

	if locs[0].Name != "wpt-001" {
		t.Errorf("Name did not match")
	}

	if locs[0].Browser != "Chrome" {
		t.Errorf("Browser did not match")
	}

	if locs[0].Label != "uswest - Agent 001" {
		t.Errorf("Label did not match")
	}

	if locs[0].Total != 1 {
		t.Errorf("Total did not match")
	}

	if locs[0].Testing != 1 {
		t.Errorf("Testing did not match")
	}

	if !locs[0].Busy {
		t.Errorf("Busy is wrong")
	}

	if locs[4].Name != "wpt-002" {
		t.Errorf("Name did not match")
	}

	if locs[4].Busy {
		t.Errorf("Idle agent is wrong")
	}

}
