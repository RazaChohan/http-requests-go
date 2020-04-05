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
	fmt.Println(*threads)
	//Loop on urls passed
	for _, urlString := range urls {
		urlString := AddSchemeToUrl(urlString)
		responseBody, err  := SendHttpRequest(urlString)
		var outputToPrint string
		if err == nil {
			outputToPrint = urlString + " << RESPONSE HASH >> " + GetMd5Hash(responseBody)
		} else {
			outputToPrint = urlString + " << RESPONSE ERROR MSG >> " + err.Error()
		}
		fmt.Println(outputToPrint)
	}
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