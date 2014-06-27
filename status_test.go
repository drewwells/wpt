package wpt

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"testing"
)

//import "github.com/kr/pretty"

func TestGetStatus(t *testing.T) {

	var testData map[string]interface{}
	//Reading Test Data
	bytes, err := ioutil.ReadFile("./test/status.json")

	if err != nil {
		t.Errorf("%v", err)
	}

	_ = json.Unmarshal(bytes, &testData)

	status := processStatus(json.Marshal(testData["testStatusComplete"]))
	t.Log("Checking status of Completed Test")

	if status.StatusCode != 200 {
		t.Errorf("StatusCode not 200")
	}
	if status.StatusText != "Test Complete" {
		t.Errorf("StatusText not 'Text Complete'")
	}

	status = processStatus(json.Marshal(testData["testStatusPending"]))
	t.Log("Checking status of Pending Test")
	if status.StatusCode != 101 {
		t.Errorf("StatusCode not 101")
	}
	if status.StatusText != "Test Pending" {
		t.Errorf("StatusText not 'Text Pending'")
	}

	status = processStatus(json.Marshal(testData["testStatusRunning"]))
	t.Log("Checking status of Running Test")
	if status.StatusCode != 100 {
		t.Errorf("StatusCode not 100")
	}
	re := regexp.MustCompile("Test Started")
	if re.FindString(status.StatusText) != "Test Started" {
		t.Errorf("StatusText not 'Text Started'")
	}

	t.Log("Checking status of invalid TestID")
	status = processStatus(json.Marshal(testData["testStatusNotFound"]))

	if status.StatusCode != 400 {
		t.Errorf("StatusCode not 400")
	}

	if status.StatusText != "Test not found" {
		t.Errorf("StatusText not 'Test not found'")
	}

}
