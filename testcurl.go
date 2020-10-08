// testcurl.go
package main

import (
	errs "errorshandler"
	"fmt"
	curl "gocurl"
	"os"
	//"time"
	//"logger"
)

var (
	//url string = "http://allprint.local"
	//url string = "https://randomnumbers.ru/generator-anglijskikh-slov"
	url = "https://httpbin.org/"
)

func main() {

	fmt.Println("Hello World!")
	options := curl.GetOptions()
	options.Url = url
	options.SetTimeout("3s")
	options.CookieFile = "/home/robot/all.cookie"
	options.FollowLocation = false
	testGetMethod(options)
}

func makeRequest(options *curl.RequestOptions) {
	client := curl.InitCurl(options)
	result := client.DoRequest()
	writeHtmlToFile(result)
}

func testGetMethod(options *curl.RequestOptions) {
	options.Url = url + "get"
	options.AddQueryParam("id", "myid")
	options.AddQueryParam("name", "Kate")
	options.AddQueryParam("action", "fuck")
	makeRequest(options)
}

func writeHtmlToFile(html string) {
	fileHandler, err := os.OpenFile("/home/robot/response.html", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	errs.ErrorHandle(err)
	defer fileHandler.Close()
	fileHandler.Truncate(0)
	fileHandler.WriteString(html)
	length := len(html)
	fmt.Println("done", length)
}
