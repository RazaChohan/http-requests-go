package main

import (
	"strconv"
	"testing"
)

//Test min function
func TestMin(t *testing.T) {
	x := min(2, 3)
	if x != 2 {
		t.Error("Expected", 5, "Got", x)
	}
}
//Test add scheme to url method
func TestAddSchemeToUrl(t *testing.T) {
	originalUrl := "www.google.com"
	url := AddSchemeToUrl(originalUrl)
	if url != "http://" + originalUrl {
		t.Error("Expected", "http://" + originalUrl, "Got", url)
	}
}
// Test do not add scheme to url if scheme already exists
func TestExistingSchemeToUrl(t *testing.T) {
	originalUrl := "http://www.google.com"
	url := AddSchemeToUrl(originalUrl)
	if url != originalUrl {
		t.Error("Expected", originalUrl, "Got", url)
	}
}
// Test md5 hash function
func TestGetMd5Hash(t *testing.T) {
	expectedHash := "774fa93897af7c00741b04ab3094175b"
	md5hash := GetMd5Hash("my_http_md5_test")
	if md5hash != expectedHash {
		t.Error("Expected", expectedHash, "Got", md5hash)
	}
}
// Test tasks per thread
func TestTaskPerThread(t *testing.T) {
	expectedTasks := 5
	tasksPerThread := TaskPerThread(25, 5)
	if expectedTasks != tasksPerThread {
		t.Error("Expected", expectedTasks, "Got", tasksPerThread)
	}
}
//Test get start and end index for
func TestGetStartAndEndIndexForFourthThread(t *testing.T) {
	startIndex, endIndex := GetStartAndEndIndexForCurrentThread(5, 4, 25)
	if startIndex != 20 || endIndex != 25 {
		t.Error("Expected", "20, 25", "Got", strconv.Itoa(startIndex) + ", " + strconv.Itoa(endIndex))
	}
}
//Test get start and end index for
func TestGetStartAndEndIndexForLastThread(t *testing.T) {
	startIndex, endIndex := GetStartAndEndIndexForCurrentThread(5, 5, 25)
	if startIndex != 25 || endIndex != 25 {
		t.Error("Expected", "25, 25", "Got", strconv.Itoa(startIndex) + ", " + strconv.Itoa(endIndex))
	}
}
// Test send http request
func TestSendHttpRequest(t *testing.T) {
	response, _ := SendHttpRequest("http://google.com")
	if len(response) == 0 {
		t.Error("Expected", "Non empty response", "Got", "Empty response")
	}
}