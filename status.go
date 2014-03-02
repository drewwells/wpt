package wpt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

type Status struct {
	StatusCode int32
	StatusText string
	Data       StatusData
}

func ProcessStatus(response []byte, err error) (status Status) {

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(response, &status)

	if err != nil {
		log.Fatal(err)
	}

	return status
}

func GetStatus(url string, key string) (status Status) {

	res, err := http.Get(url + "/testStatus.php?test=" + key)
	//res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(res.Body)

	status = ProcessStatus(bytes, err)

	//loc, _ := time.LoadLocation("America/Chicago")
	//layout := "01/02/40 10:01:01"
	//status.Data.PCompleteTime, err =
	//	time.ParseInLocation(layout, status.Data.CompleteTime, loc)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return
}
