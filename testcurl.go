// testcurl.go
package main

import (
	errs "errorshandler"
	"fmt"
	curl "gocurl"
	"math/rand"
	"os"
	"time"
)

var (
	url string = "https://www.123.ru"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello World!")
	options := curl.GetOptions()
	options.Url = url
	options.CookieFile = "/home/robot/e2.cookie"
	client := curl.InitCurl(options)
	//fmt.Println(client)
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
