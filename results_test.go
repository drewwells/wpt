package wpt

import "testing"

//import "github.com/kr/pretty"

const wpturl = "http://www.webpagetest.org"

func TestGetResult(t *testing.T) {
	t.Log("Processing Successful Result")
	response := ProcessResult([]byte(testResultSuccess))
	_ = response
	if response.Summary !=
		"http://www.webpagetest.org/results.php?test=140222_ZC_4Y9" {
		t.Errorf("Error processing Result")
	}

}
