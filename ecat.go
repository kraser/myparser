// ecat
package main

import (
	"fmt"
	"os"

	errs "github.com/kraser/errorshandler"

	curl "bitbucket.org/kravalsergey/gocurl"
	//"time"
	//"logger"
	//"net/http"

	goquery "github.com/PuerkitoBio/goquery"
)

var (
	//url string = "http://allprint.local"
	//url string = "https://randomnumbers.ru/generator-anglijskikh-slov"
	url = "https://www.e-katalog.ru"
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
	/*
		options.Url = url + "get"
		options.AddQueryParam("id", "myid")
		options.AddQueryParam("name", "Kate")
		options.AddQueryParam("action", "fuck")
	*/
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
	parse(html)
}

func parse(html string) {
	nodes, err := goquery.NewDocument(html)
	errs.ErrorHandle(err)
	nodes.Find("ul.mainmenu-list").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a.mainmenu-link").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}
