package main

import (
	"fmt"

	"github.com/kr/pretty"
)

var wpturl = "http://webpagetest.eng.wsm.local/"

func main() {
	//location := GetLocations(wpturl + "getLocations.php?f=json")
	//fmt.Printf("%# v", pretty.Formatter(location))
	results := GetResult(wpturl, "140213_GA_1D8")
	fmt.Printf("%# v", pretty.Formatter(results))
}
