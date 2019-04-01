// testcurl.go
package main

import (
	"fmt"
	curl "gocurl"
	"math/rand"
	"time"
)

var (
	url string = "http://www.dns-shop.ru"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello World!")
	options := curl.GetOptions()
	options.Url = url
	options.CookieFile = "/home/robot/e2.cookie"
	client := curl.InitCurl(options)
	fmt.Println(client)
	result := client.DoRequest(url)
	length := len(result)
	fmt.Println("done", length)

}
