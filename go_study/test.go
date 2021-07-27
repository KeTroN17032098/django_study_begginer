package main

import (
	"fmt"
	"goquery"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fedesog/webdriver"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main(url string, drivepath string) {
	chromedriver := webdriver.NewChromeDriver(drivepath)
	driverRunErr := chromedriver.Start()
	checkErr(driverRunErr)

	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}

	session, err := chromedriver.NewSession(desired, required)
	checkErr(err)

	openURLerr := session.Url(url)
	checkErr(openURLerr)

	resp, err := session.Source()
	checkErr(err)

	htmlNode, err := html.Parse(strings.NewReader(resp))
	checkErr(err)

	doc = goquery.NewDocumentFromNode(htmlNode)
	fmt.Print(doc)

}
