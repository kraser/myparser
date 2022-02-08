// ecat
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	errs "github.com/kraser/errorshandler"

	curl "bitbucket.org/kravalsergey/gocurl"
	//"time"
	//"logger"
	//"net/http"

	goquery "github.com/PuerkitoBio/goquery"
	"github.com/kraser/logger"
)

var (
	//url string = "http://allprint.local"
	//url string = "https://randomnumbers.ru/generator-anglijskikh-slov"
	url     string = "https://www.e-katalog.ru"
	logMode string
)

func init() {
	flag.StringVar(&logMode, "lm", "debug", "режим логгирования")
	//flag.StringVar(&city, "city", logMode, "город для которого разбирается прайс")

	//logMode = "debug"
}

func main() {
	flag.Parse()
	logger.SetLogLevel(logMode)
	fmt.Println("Hello World!")
	options := curl.GetOptions()
	options.Url = url
	//options.SetTimeout("3s")
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
	nodes, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	errs.ErrorHandle(err)
	nodes.Find("ul.mainmenu-list").Each(traverseTop)

}

func traverseTop(i int, s *goquery.Selection) {
	// For each item found, get the title
	anchors := s.Find("a.mainmenu-link")
	fmt.Printf("Review %d: %s\n", i, anchors)
	anchors.Each(func(i int, a *goquery.Selection) {
		href, _ := a.Attr("href")

		logger.Info(href)
		logger.Info(a.Text())
	})

}
