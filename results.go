package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type WPTResult struct {
	Run                        int32   `xml:"run"`
	URL                        string  `xml:"URL"`
	LoadTime                   int32   `xml:"loadTime"`
	TTFB                       int32   `xml:"TTFB"`
	BytesOut                   int32   `xml:"bytesOut"`
	BytesOutDoc                int32   `xml:"bytesOutDoc"`
	Connections                int32   `xml:"connections"`
	Requests                   int32   `xml:"requests"`
	RequestsDoc                int32   `xml:"requestsDoc"`
	Responses_200              int32   `xml:"responses_200"`
	Responses_404              int32   `xml:"responses_404"`
	Responses_other            int32   `xml:"responses_other"`
	Result                     int32   `xml:"result"`
	Render                     int32   `xml:"render"`
	FullyLoaded                int32   `xml:"fullyLoaded"`
	Cached                     int32   `xml:"cached"`
	DocTime                    int32   `xml:"docTime"`
	DomTime                    int32   `xml:"domTime"`
	Score_cache                int32   `xml:"score_cache"`
	Score_cdn                  int32   `xml:"score_cdn"`
	Score_gzip                 int32   `xml:"score_gzip"`
	Score_cookies              int32   `xml:"score_cookies"`
	Score_keep_alive           int32   `xml:"score_keep-alive"`
	Score_minify               int32   `xml:"score_minify"`
	Score_combine              int32   `xml:"score_combine"`
	Score_compress             int32   `xml:"score_compress"`
	Score_etags                int32   `xml:"score_etags"`
	Gzip_total                 int32   `xml:"gzip_total"`
	Gzip_savings               int32   `xml:"gzip_savings"`
	Minify_total               int32   `xml:"minify_total"`
	Minify_savings             int32   `xml:"minify_savings"`
	Image_total                int32   `xml:"image_total"`
	Image_savings              int32   `xml:"image_savings"`
	Optimization_checked       int32   `xml:"optimization_checked"`
	Aft                        int32   `xml:"aft"`
	DomElements                int32   `xml:"domElements"`
	PageSpeedVersion           float32 `xml:"pageSpeedVersion"`
	Title                      string  `xml:"title"`
	TitleTime                  int32   `xml:"titleTime"`
	LoadEventStart             int32   `xml:"loadEventStart"`
	LoadEventEnd               int32   `xml:"loadEventEnd"`
	DomContentLoadedEventStart int32   `xml:"domContentLoadedEventStart"`
	DomContentLoadedEventEnd   int32   `xml:"domContentLoadedEventEnd"`
	LastVisualChange           int32   `xml:"lastVisualChange"`
	Browser_name               string  `xml:"browser_name"`
	Browser_version            string  `xml:"browser_version"`
	Server_count               int32   `xml:"server_count"`
	Server_rtt                 int32   `xml:"server_rtt"`
	Base_page_cdn              string  `xml:"base_page_cdn"`
	Adult_site                 string  `xml:"adult_site"`
	Fixed_viewport             string  `xml:"fixed_viewport"`
	Score_progressive_jpeg     int32   `xml:"score_progressive_jpeg"`
	FirstPaint                 int32   `xml:"firstPaint"`
	DocCPUms                   float32 `xml:"docCPUms"`
	FullyLoadedCPUms           float32 `xml:"fullyLoadedCPUms"`
	DocCPUpct                  float32 `xml:"docCPUpct"`
	FullyLoadedCPUpct          float32 `xml:"fullyLoadedCPUpct"`
	Date                       int32   `xml:"date"`
	SpeedIndex                 int32   `xml:"SpeedIndex"`
	VisualComplete             int32   `xml:"visualComplete"`
	//Only available on aggregate runs
	//EffectiveBps int32 `xml:"effectiveBps"`
	//EffectiveBpsDoc int32 `xml:"effectiveBpsDoc"`
	//AvgRun int32 `xml:"avgRun"`
	//userTime dynamic values
}

type Pages struct {
	Details    string `xml:"details"`
	Checklist  string `xml:"checklist"`
	Breakdown  string `xml:"breakdown"`
	Domains    string `xml:"domains"`
	ScreenShot string `xml:"screenShot"`
}

type Thumbnails struct {
	WaterFall  string `xml:"waterfall"`
	Checklist  string `xml:"checklist"`
	ScreenShot string `xml:"screenShot"`
}

type Images struct {
	Thumbnails
	ConnectionView string `xml:"connectionView"`
}

type RawData struct {
	Headers      string `xml:"headers"`
	PageData     string `xml:"pageData"`
	RequestsData string `xml:"requestsData"`
	Utilization  string `xml:"utilization"`
}

type Frame struct {
	Time             int32  `xml:"time"`
	Image            string `xml:"image"`
	VisuallyComplete int32
}

type VideoFrames struct {
	Frame Frame `xml:"frame"`
}

type WPTResultSet struct {
	Results     WPTResult   `xml:"results"`
	Pages       Pages       `xml:"pages"`
	Thumbnails  Thumbnails  `xml:"thumbnails"`
	Images      Images      `xml:"images"`
	RawData     RawData     `xml:"rawData"`
	VideoFrames VideoFrames `xml:"videoFrames"`
}

type Views struct {
	FirstView  WPTResult `xml:"firstView"`
	RepeatView WPTResult `xml:"repeatView"`
}

type WPTRun struct {
	FirstView  WPTResultSet `xml:"firstView"`
	RepeatView WPTResultSet `xml:"repeatView"`
	Id         int32        `xml:"id"`
}

type WPTResultData struct {
	TestId            string   `xml:"testId"`
	Summary           string   `xml:"summary"`
	Location          string   `xml:"location"`
	Connectivity      string   `xml:"connectivity"`
	BwDown            int32    `xml:"bwDown"`
	BwUp              int32    `xml:"bwUp"`
	Latency           int32    `xml:"latency"`
	Plr               int32    `xml:"plr"`
	Completed         string   `xml:"completed"`
	Runs              int32    `xml:"runs"`
	SuccessfulFVRuns  int32    `xml:"successfulFVRuns"`
	Average           Views    `xml:"average"`
	Median            Views    `xml:"median"`
	StandardDeviation Views    `xml:"standardDeviation"`
	Run               []WPTRun `xml:"run"`
}

type WPTContainer struct {
	//Response   xml.Name //WPTResultResponse `xml:"response"`
	StatusCode int32         `xml:"statusCode"`
	StatusText string        `xml:"statusText"`
	Data       WPTResultData `xml:"data"`
}

func GetResult(url string, key string) (result WPTContainer) {
	res, err := http.Get(wpturl + "xmlResult/" + key + "/?f=json")
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("%+v", string(robots))
	_ = xml.Unmarshal(robots, &result)
	return
}
