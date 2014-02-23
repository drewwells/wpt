package wpt

import "testing"

//const wpturl = "http://www.webpagetest.org/"

//testResultNotFound
//testResultSuccess

func TestGetResult(t *testing.T) {
	const wpturl = "http://www.webpagetest.org/"
	_ = ProcessResult([]byte(testResultSuccess))

	response := ProcessResult([]byte(testResultSuccess))
	if response.StatusCode != 200 {
		t.Errorf("Invalid status code: %v", response.StatusCode)
	}

}
