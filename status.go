package wpt

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type StatusData struct {
	StatusCode      int32  `xml:"statusCode"`
	StatusText      string `xml:"statusText"`
	TestId          string `xml:"testId"`
	Runs            int32  `xml:"runs"`
	Fvonly          int32  `xml:"fvonly"`
	Remote          string `xml:"remote"`
	TestsExpected   int32  `xml:"testsExpected"`
	Location        string `xml:"location"`
	StartTime       string `xml:"startTime"`
	Elapsed         int32  `xml:"elapsed"`
	CompleteTime    string `xml:"completeTime"`
	PCompleteTime   time.Time
	TestsCompleted  int32 `xml:"testsCompleted"`
	FvRunsCompleted int32 `xml:"fvRunsCompleted"`
	RvRunsCompleted int32 `xml:"rvRunsCompleted"`
}

type Status struct {
	StatusCode int32      `xml:"statusCode"`
	StatusText string     `xml:"statusText"`
	Data       StatusData `xml:"data"`
}

func GetStatus(url string, key string) (status Status) {

	res, err := http.Get(url + "testStatus.php?f=xml&test=" + key)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", string(robots))
	_ = xml.Unmarshal(robots, &status)

	//loc, _ := time.LoadLocation("America/Chicago")
	//layout := "01/02/40 10:01:01"
	//status.Data.PCompleteTime, err =
	//	time.ParseInLocation(layout, status.Data.CompleteTime, loc)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return
}
