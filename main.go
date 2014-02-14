package main

import (
	"fmt"

	"github.com/kr/pretty"
)

var wpturl = "http://webpagetest.eng.wsm.local/"

func main() {
	//location := GetLocations(wpturl + "getLocations.php?f=json")
	//fmt.Printf("%# v", pretty.Formatter(location))
	results := GetResult(wpturl, "140214_XW_DS")
	fmt.Printf("%# v", pretty.Formatter(results.Data.Run[1:]))
}
