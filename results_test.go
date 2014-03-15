package wpt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

//import "github.com/kr/pretty"

//const wpturl = "http://www.webpagetest.org"
const wpturl = "http://webpagetest.eng.wsm.local"

func TestGetResult(t *testing.T) {
	//Disable to make unit testing faster
	return
	_, err := GetResult("http://notaurl", "140314_7Q_A")

	if err == nil {
		t.Errorf("No error thrown for non-existant url")
		return
	}
	re, _ := regexp.Compile(`dial tcp`)

	if re.MatchString(fmt.Sprintf("%v", err)) != true {
		t.Errorf("%v", err)
	}

}

func TestProcessResult(t *testing.T) {

	t.Log("Processing Successful Result")

	var testData map[string]interface{}

	bytes, err := ioutil.ReadFile("./test/result.json")

	if err != nil {
		t.Errorf("%v", err)
	}

	_ = json.Unmarshal(bytes, &testData)

	//testResultWaiting
	result, _ := ProcessResult(json.Marshal(testData["testResultWaiting"]))

	if result.StatusCode != 101 {
		t.Errorf("StatusCode not 101")
	}

	if result.StatusText != "Waiting behind 1 other test..." {
		t.Errorf("Improper parsing of waiting status text")
	}

	//testResultFront
	result, _ = ProcessResult(json.Marshal(testData["testResultFront"]))

	if result.StatusCode != 101 {
		t.Errorf("StatusCode not 101")
	}

	if result.StatusText != "Waiting at the front of the queue..." {
		t.Errorf("Improper parsing of waiting at front status text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultRunning
	result, _ = ProcessResult(json.Marshal(testData["testResultRunning"]))

	if result.StatusCode != 100 {
		t.Errorf("StatusCode not 100")
	}

	if result.StatusText != "Test Started 2 seconds ago" {
		t.Errorf("Improper parsing of waiting of running... text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultNotFound
	result, _ = ProcessResult(json.Marshal(testData["testResultNotFound"]))

	if result.StatusCode != 400 {
		t.Errorf("StatusCode not 400")
	}

	if result.StatusText != "Test not found" {
		t.Errorf("Improper parsing of status text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultSuccess
	result, _ = ProcessResult(json.Marshal(testData["testResultSuccess"]))

	if result.Data.Url != "http://www.123greetings.com/birthday/happy_birthday/birthday162.html" {
		t.Errorf("Invalid URL")
		t.Errorf("Found: " + result.Data.Url)
	}

	if result.Data.Summary !=
		"http://www.webpagetest.org/results.php?test=140222_ZC_4Y9" {
		t.Errorf("Error processing Result")
	}

	if result.Data.Runs[0].FirstView.TTFB != 690 {
		t.Errorf("TTFB in for first Run invalid")
	}

	if result.Data.Completed != 1393047807 {
		t.Errorf("Completed timestamp invalid")
	}

}
