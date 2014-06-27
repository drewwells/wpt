package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Status(url string, key string) PStatus {

	res, err := http.Get(url + "/testStatus.php?test=" + key)
	//res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(res.Body)

	return processStatus(bytes, err)
}

func processStatus(response []byte, err error) (status PStatus) {

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(response, &status)

	if err != nil {
		log.Fatal(err)
	}

	return status
}

type StatusData struct {
	StatusCode      int32
	StatusText      string
	TestId          string
	Runs            int32
	Fvonly          int32
	Remote          bool
	TestsExpected   int32
	Location        string
	StartTime       string
	Elapsed         int32
	CompleteTime    string
	TestsCompleted  int32
	FvRunsCompleted int32
	RvRunsCompleted int32
}

type PStatus struct {
	StatusCode int32
	StatusText string
	Data       StatusData
}
