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
	"strings"

	"labix.org/v2/mgo/bson"
	//"labix.org/v2/mgo/bson"
	//"github.com/kr/pretty"
)

func process(response []byte, err error) (Result, error) {
	var (
		jsonR  ResultJSON
		result Result
		jsonFR ResultFJSON
	)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &jsonR)

	//Handle inconsistent return from running test (where runs is a number)
	if err != nil {

		err = json.Unmarshal(response, &jsonFR)
		if err != nil {
			log.Fatal("%+v", err)
		}
		result.StatusCode = jsonFR.StatusCode
		result.StatusText = jsonFR.StatusText
	} else {
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
	}

	return result, err
}

func Get(url string, key string) (Result, error) {

	res, err := http.Get(url + "/jsonResult.php?test=" + key)

	if err != nil {
		return Result{}, err
	}

	defer res.Body.Close()
	return process(ioutil.ReadAll(res.Body))
}

type WPTResult struct {
	Run                        int32   `json:"run"`
	URL                        string  `json:"URL"`
	LoadTime                   int32   `json:"loadTime"`
	TTFB                       int32   `json:"TTFB"`
	BytesIn                    float32 `json:"bytesIn"`
	BytesInDoc                 int32
	BytesOut                   int32                  `json:"bytesOut"`
	BytesOutDoc                int32                  `json:"bytesOutDoc"`
	Connections                int32                  `json:"connections"`
	Requests                   []interface{}          `json:"requests"`
	RequestsDoc                int32                  `json:"requestsDoc"`
	Responses_200              int32                  `json:"responses_200"`
	Responses_404              int32                  `json:"responses_404"`
	Responses_other            int32                  `json:"responses_other"`
	Result                     int32                  `json:"result"`
	Render                     int32                  `json:"render"`
	FullyLoaded                int32                  `json:"fullyLoaded"`
	Cached                     int32                  `json:"cached"`
	DocTime                    int32                  `json:"docTime"`
	DomTime                    int32                  `json:"domTime"`
	Score_cache                int32                  `json:"score_cache"`
	Score_cdn                  int32                  `json:"score_cdn"`
	Score_gzip                 int32                  `json:"score_gzip"`
	Score_cookies              int32                  `json:"score_cookies"`
	Score_keep_alive           int32                  `json:"score_keep-alive"`
	Score_minify               int32                  `json:"score_minify"`
	Score_combine              int32                  `json:"score_combine"`
	Score_compress             int32                  `json:"score_compress"`
	Score_etags                int32                  `json:"score_etags"`
	Gzip_total                 float64                `json:"gzip_total"`
	Gzip_savings               int32                  `json:"gzip_savings"`
	Minify_total               int32                  `json:"minify_total"`
	Minify_savings             int32                  `json:"minify_savings"`
	Image_total                int32                  `json:"image_total"`
	Image_savings              int32                  `json:"image_savings"`
	Optimization_checked       int32                  `json:"optimization_checked"`
	Aft                        int32                  `json:"aft"`
	DomElements                int32                  `json:"domElements"`
	PageSpeedVersion           float32                `json:"pageSpeedVersion,string"`
	Title                      string                 `json:"title"`
	TitleTime                  int32                  `json:"titleTime"`
	LoadEventStart             int32                  `json:"loadEventStart"`
	LoadEventEnd               int32                  `json:"loadEventEnd"`
	DomContentLoadedEventStart int32                  `json:"domContentLoadedEventStart"`
	DomContentLoadedEventEnd   int32                  `json:"domContentLoadedEventEnd"`
	LastVisualChange           int32                  `json:"lastVisualChange"`
	Browser_Name               string                 `json:"browser_name"`
	Browser_Version            string                 `json:"browser_version"`
	Server_count               int32                  `json:"server_count"`
	Server_rtt                 int32                  `json:"server_rtt"`
	Base_page_cdn              string                 `json:"base_page_cdn"`
	Adult_site                 int32                  `json:"adult_site"`
	Fixed_viewport             int32                  `json:"fixed_viewport"`
	Score_progressive_jpeg     int32                  `json:"score_progressive_jpeg"`
	FirstPaint                 int32                  `json:"firstPaint"`
	DocCPUms                   float32                `json:"docCPUms"`
	FullyLoadedCPUms           float32                `json:"fullyLoadedCPUms"`
	DocCPUpct                  float32                `json:"docCPUpct"`
	FullyLoadedCPUpct          float32                `json:"fullyLoadedCPUpct"`
	Date                       float64                `json:"date"`
	SpeedIndex                 int32                  `json:"SpeedIndex"`
	VisualComplete             int32                  `json:"visualComplete"`
	UserTiming                 map[string]int         `json:"-" bson:"usertime"`
	Extra                      map[string]interface{} `json:"-" bson:"_notincluded,inline"`
}

func (t *WPTResult) MarshalJSON() ([]byte, error) {
	var j interface{}
	b, _ := bson.Marshal(t)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

func (t *WPTResult) UnmarshalJSON(b []byte) error {
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	//Delete erroneous data value
	delete(j, "userTime")
	for key, val := range j {
		j[strings.ToLower(key)] = val
	}
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, t)
}

type Pages struct {
	Details    string `json:"details"`
	Checklist  string `json:"checklist"`
	Breakdown  string `json:"breakdown"`
	Domains    string `json:"domains"`
	ScreenShot string `json:"screenShot"`
}

type Thumbnails struct {
	WaterFall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}

type Images struct {
	Thumbnails     `bson:",inline"`
	ConnectionView string `json:"connectionView"`
}

type RawData struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
	Utilization  string `json:"utilization"`
}

type VideoFrame struct {
	Time             int32  `json:"time"`
	Image            string `json:"image"`
	VisuallyComplete int32
}

type WPTResultSet struct {
	WPTResult   `bson:",inline"`
	Pages       Pages        `json:"pages"`
	Thumbnails  Thumbnails   `json:"thumbnails"`
	Images      Images       `json:"images"`
	RawData     RawData      `json:"rawData"`
	VideoFrames []VideoFrame `json:"videoFrames"`
}

type Views struct {
	FirstView  WPTResult //`json:"firstView"`
	RepeatView WPTResult //`json:"repeatView"`
}

type WPTRun struct {
	FirstView WPTResultSet `json:"firstView"`
	//RepeatView WPTResultSet `json:"repeatView"`
	Id int32 `json:"id"`
}

type WPTBaseResultData struct {
	Url              string  `json:"url" bson:"testurl"`
	TestId           string  `json:"testId" bson:"testid"`
	Summary          string  `json:"summary"`
	Location         string  `json:"location"`
	From             string  `json:"from"`
	Label            string  `json:"label" bson:"label"`
	Connectivity     string  `json:"connectivity"`
	BwDown           int32   `json:"bwDown"`
	BwUp             int32   `json:"bwUp"`
	Latency          int32   `json:"latency"`
	Completed        float64 `json:"completed"`
	SuccessfulFVRuns int32   `json:"successfulFVRuns"`
	//Average           Views  `json:"average"`
	//Median            Views  `json:"median"`
	//StandardDeviation Views  `json:"standardDeviation"`
}

type WPTResultRawData struct {
	WPTBaseResultData
	TestId     string `json:"id"`
	StatusCode int32
	StatusText string
	Plr        interface{}       `json:"plr"`
	Runs       map[string]WPTRun `json:"runs"`
}

type ResultJSON struct {
	StatusCode int32            `json:"statusCode"`
	StatusText string           `json:"statusText"`
	Completed  float64          `json:"completed"`
	Data       WPTResultRawData `json:"data" bson:"data"`
}

type WPTResultCleanData struct {
	WPTBaseResultData `bson:",inline"`
	Plr               int32    `bson:",minsize"`
	Runs              []WPTRun `json:"runs" bson:"runs"`
}

//Special struct to handle unstructured wpt responses
type ResultFJSON struct {
	StatusCode int32   `json:"statusCode"`
	StatusText string  `json:"statusText"`
	Completed  float64 `json:"completed"`
}

type Result struct {
	StatusCode int32
	StatusText string
	Data       WPTResultCleanData `json:"data"`
}
