package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type WPTResult struct {
	URL                        string  `xml:"URL"`
	LoadTime                   int32   `xml:"loadTime"`
	TTFB                       int32   `xml:"TTFB"`
	BytesOut                   int32   `xml:"bytesOut"`
	BytesOutDoc                int32   `xml:"bytesOutDoc"`
	connections                int32   `xml:"connections"`
	requests                   int32   `xml:"requests"`
	requestsDoc                int32   `xml:"requestsDoc"`
	responses_200              int32   `xml:"responses_200"`
	responses_404              int32   `xml:"responses_404"`
	responses_other            int32   `xml:"responses_other"`
	result                     int32   `xml:"result"`
	render                     int32   `xml:"render"`
	fullyLoaded                int32   `xml:"fullyLoaded"`
	cached                     int32   `xml:"cached"`
	docTime                    int32   `xml:"docTime"`
	domTime                    int32   `xml:"domTime"`
	score_cache                int32   `xml:"score_cache"`
	score_cdn                  int32   `xml:"score_cdn"`
	score_gzip                 int32   `xml:"score_gzip"`
	score_cookies              int32   `xml:"score_cookies"`
	score_keep_alive           int32   `xml:"score_keep-alive"`
	score_minify               int32   `xml:"score_minify"`
	score_combine              int32   `xml:"score_combine"`
	score_compress             int32   `xml:"score_compress"`
	score_etags                int32   `xml:"score_etags"`
	gzip_total                 int32   `xml:"gzip_total"`
	gzip_savings               int32   `xml:"gzip_savings"`
	minify_total               int32   `xml:"minify_total"`
	minify_savings             int32   `xml:"minify_savings"`
	image_total                int32   `xml:"image_total"`
	image_savings              int32   `xml:"image_savings"`
	optimization_checked       int32   `xml:"optimization_checked"`
	aft                        int32   `xml:"aft"`
	domElements                int32   `xml:"domElements"`
	pageSpeedVersion           int32   `xml:"pageSpeedVersion"`
	title                      int32   `xml:"title"`
	titleTime                  int32   `xml:"titleTime"`
	loadEventStart             int32   `xml:"loadEventStart"`
	loadEventEnd               int32   `xml:"loadEventEnd"`
	domContentLoadedEventStart int32   `xml:"domContentLoadedEventStart"`
	domContentLoadedEventEnd   int32   `xml:"domContentLoadedEventEnd"`
	lastVisualChange           int32   `xml:"lastVisualChange"`
	browser_name               string  `xml:"browser_name"`
	browser_version            float32 `xml:"browser_version"`
	server_count               int32   `xml:"server_count"`
	server_rtt                 int32   `xml:"server_rtt"`
	base_page_cdn              string  `xml:"base_page_cdn"`
	adult_site                 string  `xml:"adult_site"`
	fixed_viewport             string  `xml:"fixed_viewport"`
	score_progressive_jpeg     int32   `xml:"score_progressive_jpeg"`
	firstPaint                 int32   `xml:"firstPaint"`
	docCPUms                   int32   `xml:"docCPUms"`
	fullyLoadedCPUms           int32   `xml:"fullyLoadedCPUms"`
	docCPUpct                  int32   `xml:"docCPUpct"`
	fullyLoadedCPUpct          int32   `xml:"fullyLoadedCPUpct"`
	date                       int32   `xml:"date"`
	SpeedIndex                 int32   `xml:"SpeedIndex"`
	visualComplete             int32   `xml:"visualComplete"`
	effectiveBps               int32   `xml:"effectiveBps"`
	effectiveBpsDoc            int32   `xml:"effectiveBpsDoc"`
	avgRun                     int32   `xml:"avgRun"`
	//userTime dynamic values
}

type FirstView struct {
	FirstView WPTResult `xml:"firstView"`
}

type WPTRun struct {
	FirstView WPTResult `xml:"firstView>results"`
	Id        int32     `xml:"id"`
}

type WPTResultData struct {
	TestId            string    `xml:"testId"`
	Summary           string    `xml:"summary"`
	Location          string    `xml:"location"`
	Connectivity      string    `xml:"connectivity"`
	BwDown            int32     `xml:"bwDown"`
	BwUp              int32     `xml:"bwUp"`
	Latency           int32     `xml:"latency"`
	Plr               int32     `xml:"plr"`
	Completed         string    `xml:"completed"`
	Runs              int32     `xml:"runs"`
	SuccessfulFVRuns  int32     `xml:"successfulFVRuns"`
	Average           FirstView `xml:"average"`
	Median            FirstView `xml:"median"`
	StandardDeviation FirstView `xml:"standardDeviation"`
	Run               []WPTRun  `xml:"run"`
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
