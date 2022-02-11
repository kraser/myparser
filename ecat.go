// ecat
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	errs "github.com/kraser/errorshandler"
	price "github.com/kraser/goprice"

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
	url             string = "https://www.e-katalog.ru"
	logMode         string
	myprice         *price.Price
	currentCategory *price.Category
)

func init() {
	flag.StringVar(&logMode, "lm", "info", "режим логгирования")
	//flag.StringVar(&city, "city", logMode, "город для которого разбирается прайс")

	//logMode = "debug"
}

func main() {
	flag.Parse()
	logger.SetLogLevel(logMode)
	fmt.Println("Hello World!")
	myprice = price.GetPrice("e-catalog")
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
	fmt.Println(myprice)
}

func parse(html string) {
	nodes, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	errs.ErrorHandle(err)
	nodes.Find("ul.mainmenu-list li.mainmenu-item").Each(traverseTop)

}

func traverseTop(i int, s *goquery.Selection) {
	//logger.Info(s.Children().Length())
	// For each item found, get the title
	anchors := s.Find("a.mainmenu-link")
	anchors.Each(func(i int, a *goquery.Selection) {
		category := price.CreateCategory(a.Text())
		href, _ := a.Attr("href")
		logger.Info(a.Text(), href)
		category.SetUrl(href)
		myprice.PushCategory(category)
		div := s.Find("div.mainmenu-sublist").First()
		div.Find("a.mainmenu-subitem").Each(traverseSubCat)
		//category = currentCategory
		myprice.AddCategory(category)
	})

}

func traverseSubCat(i int, s *goquery.Selection) {
	//logger.Info(s.Children().Length())
	href, _ := s.Attr("href")
	var name = strings.Trim(s.Text(), " \n")
	logger.Info(name, href)
	category := price.CreateCategory(name)
	category.SetUrl(href)
	currentCategory = myprice.GetCurrentCategory()
	logger.Info(currentCategory)
	currentCategory.AddCategory(category)
	//myprice.AddCategory(category)
	/*
		// For each item found, get the title
		anchors := s.Find("a.mainmenu-link")
		anchors.Each(func(i int, a *goquery.Selection) {
			category := price.CreateCategory(a.Text())
			href, _ := a.Attr("href")
			logger.Info(a.Text(), href)
			category.SetUrl(href)
			myprice.AddCategory(category)
			div := s.Find("div.mainmenu-sublist").First()
			div.Find("a.mainmenu-subitem").Each(traverseSubCat)
		})
	*/
}
