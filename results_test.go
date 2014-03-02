package wpt

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

//import "github.com/kr/pretty"

const wpturl = "http://www.webpagetest.org"

func TestGetResult(t *testing.T) {
	t.Log("Processing Successful Result")

	var testData map[string]interface{}

	bytes, err := ioutil.ReadFile("./test/result.json")

	if err != nil {
		t.Errorf("%v", err)
	}

	_ = json.Unmarshal(bytes, &testData)

	//testResultNotFound
	//testResultWaiting
	//testResultFront
	//testResultRunning

	result := ProcessResult(json.Marshal(testData["testResultSuccess"]))

	if result.Summary !=
		"http://www.webpagetest.org/results.php?test=140222_ZC_4Y9" {
		t.Errorf("Error processing Result")
	}

	if result.Runs[0].FirstView.TTFB != 690 {
		t.Errorf("TTFB in for first Run invalid")
	}

	if result.Completed != 1393047807 {
		t.Errorf("Completed timestamp invalid")
	}

}
