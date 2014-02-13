package main

import (
	"fmt"

	"github.com/kr/pretty"
)

var wpturl = "http://webpagetest.eng.wsm.local/getLocations.php?f=json"

func main() {
	location := GetLocations(wpturl)
	fmt.Printf("%# v", pretty.Formatter(location))
}
