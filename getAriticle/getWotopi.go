package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func GetWotopiBirth() {
	// Request the HTML page.
	res, err := http.Get("https://wotopi.jp/archives/category/birth")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	titles := getTitles(doc)
	fmt.Println(titles)
}

func getTitles(doc *goquery.Document) []string {
	var array []string
	doc.Find("#main-contents .figure-list article .content-title").Each(func(i int, s *goquery.Selection) {
		if i < 15 {
			array = append(array, s.Text())
		}
	})
	return array
}

func main() {
	GetWotopiBirth()
}
