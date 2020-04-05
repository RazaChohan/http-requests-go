package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//Get max number of threads required
	threads := flag.Int("parallel", 10, "Number of parallel threads")
	flag.Parse()
	//get all urls from flags
	urls := flag.Args()
	fmt.Println(*threads)
	//Loop on urls passed
	for _, url := range urls {
		responseBody := SendHttpRequest(url)
		fmt.Println(url + " " + GetMd5Hash(responseBody))
	}
}
//Send http request
func SendHttpRequest(url string) string {
	response, err := http.Get(url)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}
//Get MD5 hash
func GetMd5Hash(responseBody string) string {
	hash := md5.Sum([]byte(responseBody))
	return hex.EncodeToString(hash[:])

}