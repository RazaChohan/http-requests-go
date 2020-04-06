package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//Get max number of threads required
	threads := flag.Int("parallel", 10, "Number of parallel threads")
	flag.Parse()
	//get all urls from flags
	urls := flag.Args()
	//Calculate urls per threads
	threadResults := make(chan []string)
	//Number of urls
	urlsSize := len(urls)
	//If no of thread greater than urls update threads to no. of urls
	if *threads > urlsSize {
		*threads = urlsSize
	}
	//no. of tasks per thread
	taskPerThread := TaskPerThread(urlsSize, *threads)
	// Goroutine to send multiple jobs to the channel
	for i := 0; i < *threads; i++ {
		//Get start and end indexes of urls
		startIndexOfUrls, endIndexOfUrls := GetStartAndEndIndexForCurrentThread(taskPerThread, i, urlsSize)
		go func(num int) {
			threadResults <- ProcessUrlRequests(urls[startIndexOfUrls:endIndexOfUrls])
		}(i)
	}

	// Receive output from the channel to print
	for i := 0; i < *threads; i++ {
		for _, stringToPrint := range <-threadResults {
			fmt.Println(stringToPrint)
		}
	}
}
// process url requests
func ProcessUrlRequests(urls []string) []string {
	var outputsToPrint []string
	//Loop on urls passed
	for _, urlString := range urls {
		//Add scheme to url if not added
		urlString := AddSchemeToUrl(urlString)
		//Send get requests
		responseBody, err  := SendHttpRequest(urlString)
		var outputToPrint string
		if err == nil { //If no error get md5 hash of response body
			outputToPrint = urlString + " " + GetMd5Hash(responseBody)
		} else { // if error print error in output
			outputToPrint = urlString + " Error: " + err.Error()
		}
		outputsToPrint = append(outputsToPrint, outputToPrint)
	}
	return outputsToPrint
}
//Send http request
func SendHttpRequest(url string) (string, error) {
	response, err := http.Get(url)
	var responseStr string
	if err == nil && response.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		responseStr = string(bodyBytes)
	}
	return responseStr, err
}
//Get MD5 hash
func GetMd5Hash(responseBody string) string {
	hash := md5.Sum([]byte(responseBody))
	return hex.EncodeToString(hash[:])
}
//Add scheme to url
func AddSchemeToUrl(urlStr string) string {
	parsedUrl, _ := url.Parse(urlStr)
	if len(parsedUrl.Scheme) == 0 {
		urlStr = "http://" + urlStr
 	}
	return urlStr
}
//Tasks per thread
func TaskPerThread(urlsSize int, allowedThreads int) int {
	return (urlsSize + allowedThreads - 1) / allowedThreads
}
//Start and end range of element for current thread
func GetStartAndEndIndexForCurrentThread(tasksPerThread int, threadNumber int, totalTasks int) (int, int){
	startIndex := threadNumber * tasksPerThread
	endIndex := min(startIndex + tasksPerThread, totalTasks)
	return startIndex, endIndex
}
// Get min of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}