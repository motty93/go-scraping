package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	// "github.com/sclevine/agouti"
)

func scrape() {
	res, err := http.Get("https://moridaira.jp/morris/performers-edition-morris")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	selection := doc.Find("div.morrisshowbox")
	selection.Each(findElements)
}

func findElements(i int, s *goquery.Selection) {
	title := s.Text()
	url, exists := s.Find("a").Find("img").Attr("src")
	if exists != true {
		log.Fatalln("image not exists")
	}
	fmt.Printf("title: %s, image: %s\n", title, url)
}

func main() {
	// options := agouti.ChromeOptions(
	// 	"args", []string{
	// 		"--headless", "--disable-gpu",
	// 	},
	// )
	// driver := agouti.ChromeDriver(agouti.Browser("chrome"), agouti.Debug, options)
	// if err := driver.Start(); err != nil {
	// 	log.Fatalf("Failed to start driver: %v", err)
	// }
	// goroutineの形 wg = sync.WaitGroupを使う形へ
	scrape()
}
