// testcurl.go
package main

import (
	errs "errorshandler"
	"fmt"
	curl "gocurl"

	//"math/rand"
	"os"
	//"time"
	"logger"
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
	logger.Debug(options.Url)
	options.SetTimeout("3s")
	options.CookieFile = "/home/robot/all.cookie"
	options.FollowLocation = false
	options.SetMethod("post")
	client := curl.InitCurl(options)
	result := client.DoRequest(url)
	writeHtmlToFile(result)
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

//func newmain() {
//	myURL := "http://www.jonathanmh.com"
//	nextURL := myURL
//	var i int
//	for i < 100 {
//		client := &http.Client{
//			CheckRedirect: func(req *http.Request, via []*http.Request) error {
//				return http.ErrUseLastResponse
//			},
//		}
//		resp, err := client.Get(nextURL)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("StatusCode:", resp.StatusCode)
//		fmt.Println(resp.Request.URL)
//		if resp.StatusCode == 200 {
//			fmt.Println("Done!")
//			break
//		} else {
//			nextURL = resp.Header.Get("Location")
//		}
//	}
//}
