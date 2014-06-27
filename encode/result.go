package encode

import (
	"encoding/json"
	"strings"

	"labix.org/v2/mgo/bson"
)

type R struct {
	rbase `bson:",inline"`
	Plr   int32  `bson:",minsize"`
	Runs  []rrun `json:"runs" bson:"runs"`
}

type rbase struct {
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

type PResult struct {
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

func (t *PResult) MarshalJSON() ([]byte, error) {
	var j interface{}
	b, _ := bson.Marshal(t)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

func (t *PResult) UnmarshalJSON(b []byte) error {
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

type VideoFrame struct {
	Time             int32  `json:"time"`
	Image            string `json:"image"`
	VisuallyComplete int32
}

type PResultSet struct {
	PResult     `bson:",inline"`
	Pages       Pages        `json:"pages"`
	Thumbnails  Thumbnails   `json:"thumbnails"`
	Images      Images       `json:"images"`
	RawData     rawdata      `json:"rawData"`
	VideoFrames []VideoFrame `json:"videoFrames"`
}

type Views struct {
	FirstView  PResult //`json:"firstView"`
	RepeatView PResult //`json:"repeatView"`
}

type rrun struct {
	FirstView PResultSet `json:"firstView"`
	//RepeatView PResultSet `json:"repeatView"`
	Id int32 `json:"id"`
}

type rawdata struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
	Utilization  string `json:"utilization"`
}

type rrrun struct {
	rbase
	TestId     string `json:"id"`
	StatusCode int32
	StatusText string
	Plr        interface{}     `json:"plr"`
	Runs       map[string]rrun `json:"runs"`
}

type JResult struct {
	StatusCode int32   `json:"statusCode"`
	StatusText string  `json:"statusText"`
	Completed  float64 `json:"completed"`
	Data       rrrun   `json:"data" bson:"data"`
}
