/*
Package wpt provides methods and native go types for consuming
data from a WebPageTest server.  This is useful for getting
WebPageTest results and test status.


It provides a function, Get, for retrieiving all data about
a test.  Status is useful for checking the current progress
of a run.
*/
package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/drewwells/wpt/encode"
)

type Response struct {
	StatusCode int32
	StatusText string
	Data       encode.R `json:"data"`
}

// Get takes the server url and testid and returns the test
// details of a WebPageTest run.
//
// A nil object and error is returned if there are errors
// communicating with the server.
func Get(url string, key string) (Response, error) {

	res, err := http.Get(url + "/jsonResult.php?test=" + key)

	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()
	return process(ioutil.ReadAll(res.Body))
}

// Convert JSON to Go objects
// Handles inconsistencies in the WebPageTest API
func process(response []byte, err error) (Response, error) {
	var (
		jsonR  encode.ResultJSON
		result Response
		jsonFR encode.ResultFJSON
	)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &jsonR)
	//Handle inconsistent return from running test
	// (where runs is a number)
	if err != nil {

		err = json.Unmarshal(response, &jsonFR)
		if err != nil {
			log.Fatal(err)
		}
		result.StatusCode = jsonFR.StatusCode
		result.StatusText = jsonFR.StatusText
		return result, err

	}
	//Lots of work to convert {"0":{},"1":{}} to [{},{}]
	result.StatusCode = jsonR.StatusCode
	result.StatusText = jsonR.StatusText
	result.Data.TestId = jsonR.Data.TestId
	result.Data.Summary = jsonR.Data.Summary
	result.Data.Label = jsonR.Data.Label
	result.Data.Url = jsonR.Data.Url
	result.Data.Location = jsonR.Data.Location
	result.Data.Connectivity = jsonR.Data.Connectivity
	result.Data.BwDown = jsonR.Data.BwDown
	result.Data.BwUp = jsonR.Data.BwUp
	result.Data.Latency = jsonR.Data.Latency

	//iOS app sends int, but browsers send string
	switch v := jsonR.Data.Plr.(type) {
	case string:
		inf, _ := strconv.ParseInt(v, 10, 32)
		result.Data.Plr = int32(inf)
	case int64:
		result.Data.Plr = int32(v)
	case float64:
		result.Data.Plr = int32(v)
	case nil:
		result.Data.Plr = 1
	default:
		log.Printf("Failed to convert: %v", v)
	}
	result.Data.Completed = jsonR.Data.Completed
	result.Data.SuccessfulFVRuns = jsonR.Data.SuccessfulFVRuns

	r, _ := regexp.Compile("^userTime.(.*)")
	for i, val := range jsonR.Data.Runs {
		_ = i

		val.FirstView.UserTiming = make(map[string]int)

		for key, extra := range val.FirstView.Extra {
			if r.MatchString(key) && extra != nil {
				metric := r.FindStringSubmatch(key)[1]
				val.FirstView.UserTiming[metric] = int(extra.(float64))
			}
		}
		result.Data.Runs = append(result.Data.Runs, val)
	}

	return result, err
}
