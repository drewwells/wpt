package wpt

import (
	"fmt"
	"testing"
)

//import "github.com/kr/pretty"
import "github.com/kr/pretty"

const wpturl = "http://webpagetest.eng.wsm.local/"

func TestGetResult(t *testing.T) {

	response := ProcessResult([]byte(testResultSuccess))
	_ = response
	fmt.Printf("%# v", pretty.Formatter(response))

}
