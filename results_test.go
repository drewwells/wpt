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

func TestGet(t *testing.T) {
	//Disable to make unit testing faster
	return
	_, err := Get("http://notaurl", "140314_7Q_A")

	if err == nil {
		t.Errorf("No error thrown for non-existant url")
		return

	}
	re, _ := regexp.Compile(`dial tcp`)

	if re.MatchString(fmt.Sprintf("%v", err)) != true {
		t.Errorf("%v", err)
	}

}

func TestPending(t *testing.T) {

	var testData map[string]interface{}
	bytes, err := ioutil.ReadFile("./test/result.json")

	if err != nil {
		t.Errorf(err.Error())
	}
	_ = json.Unmarshal(bytes, &testData)

	//testResultWaiting
	result, _ := process(json.Marshal(testData["testResultWaiting"]))

	if result.StatusCode != 101 {
		t.Errorf("StatusCode not 101")
	}

	if result.StatusText != "Waiting behind 1 other test..." {
		t.Errorf("Improper parsing of waiting status text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultFront
	result, _ = process(json.Marshal(testData["testResultFront"]))
	if result.StatusCode != 101 {
		t.Errorf("StatusCode not 101")
	}

	if result.StatusText != "Waiting at the front of the queue..." {
		t.Errorf("Improper parsing of waiting at front status text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultRunning
	result, _ = process(json.Marshal(testData["testResultRunning"]))

	if result.StatusCode != 100 {
		t.Errorf("StatusCode not 100")
	}

	if result.StatusText != "Test Started 2 seconds ago" {
		t.Errorf("Improper parsing of waiting of running... text")
		t.Errorf("Found: " + result.StatusText)
	}

}

func TestProcess(t *testing.T) {

	t.Log("Processing Successful Result")

	var testData map[string]interface{}

	bytes, err := ioutil.ReadFile("./test/result.json")

	if err != nil {
		t.Errorf("%v", err)
	}

	_ = json.Unmarshal(bytes, &testData)

	//testResultNotFound
	result, _ := process(json.Marshal(testData["testResultNotFound"]))

	if result.StatusCode != 400 {
		t.Errorf("StatusCode not 400")
	}

	if result.StatusText != "Test not found" {
		t.Errorf("Improper parsing of status text")
		t.Errorf("Found: " + result.StatusText)
	}

	//testResultSuccess
	result, _ = process(json.Marshal(testData["testResultSuccess"]))
	//poop, _ := json.Marshal(testData["testResultSuccess"])
	//t.Errorf("%+v", string(poop))
	if result.Data.Url != "http://www.retailmenot.test-prodpreview-new/view/macys.com" {
		t.Errorf("Invalid URL")
		t.Errorf("Found: %#v Expected: %#v", result.Data.Url, "http://www.retailmenot.test-prodpreview-new/view/macys.com")
	}
	expectedSummary := "http://webpagetest.eng.wsm.local/results.php?test=140408_FA_2"
	if result.Data.Summary != expectedSummary {
		t.Errorf("Invalid summary link\n")
		t.Errorf("Found: %v Expected: %v", result.Data.Summary, expectedSummary)
	}
	conn := "DSL"
	if result.Data.Connectivity != conn {
		t.Errorf("Invalid From string\n")
		t.Errorf("\nFound: %s\nExpected: %s",
			result.Data.Connectivity,
			conn,
		)
	}
	from := " - <b>Chrome</b> - <b>DSL</b>"
	if result.Data.From != from {
		t.Errorf("Invalid From string\n")
		t.Errorf("\nFound: %s\nExpected: %s", result.Data.From, from)
	}
	expFirstOffer := 3615
	extra := result.Data.Runs[0].FirstView.Extra
	if len(result.Data.Runs) < 1 {
		t.Errorf("No runs found.")
	} else if result.Data.Runs[0].FirstView.TTFB != 1725 {
		t.Errorf("TTFB in for first Run invalid")
		t.Errorf("Found: %v expected: %v", result.Data.Runs[0].FirstView.TTFB, 1725)
	} else if extra["usertime.first_offer"] != nil && int(extra["usertime.first_offer"].(float64)) != expFirstOffer {
		t.Errorf("User timing data not found")
		t.Errorf("Found: %v, Expected: %v", extra["userTime.first_offer"], expFirstOffer)
	} else if int(result.Data.Runs[0].FirstView.UserTiming["first_offer"]) != expFirstOffer {
		t.Errorf("User timing: first_offer parsing error")
		t.Errorf("Found: %v, Expected: %v", result.Data.Runs[0].FirstView.UserTiming["first_offer"], expFirstOffer)
	} else if result.Data.Runs[0].FirstView.UserTiming["last_offer"] != 3660 {
		t.Errorf("User timing: last_offer parsing error")
		t.Errorf("Found: %v, Expected: %v", result.Data.Runs[0].FirstView.UserTiming["last_offer"], 3660)
	}

	if result.Data.Completed != 1396935254 {
		t.Errorf("Completed timestamp invalid")
		t.Errorf("Found: %v expected: %v", result.Data.Completed, 1396935254)
	}

}
