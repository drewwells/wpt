package wpt

import (
	"fmt"
	"testing"
)
import "github.com/kr/pretty"

//const wpturl = "http://www.webpagetest.org/"

//testResultNotFound
//testResultSuccess
const wpturl = "http://webpagetest.eng.wsm.local/"

func TestGetResult(t *testing.T) {

	response := ProcessResult([]byte(testResultSuccess))

	fmt.Printf("%# v", pretty.Formatter(response.Runs))

}
