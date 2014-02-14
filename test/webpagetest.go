package main

import (
	"fmt"

	wpt "github.com/drewwells/webpagetest-go"
	"github.com/kr/pretty"
)

var wpturl = "http://webpagetest.eng.wsm.local/"

func main() {
	//location := GetLocations(wpturl + "getLocations.php?f=json")
	//fmt.Printf("%# v", pretty.Formatter(location))
	//results := GetResult(wpturl, "140214_XW_DS")
	//fmt.Printf("%# v", pretty.Formatter(results))
	status := wpt.GetStatus(wpturl, "140214_XW_DS")
	fmt.Printf("%# v", pretty.Formatter(status))
}
