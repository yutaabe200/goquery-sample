package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func GetFreakShowWorks() {
	// Request the HTML page.
	res, err := http.Get("https://freak-show.jp/")
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

	doc.Find("#one section .inner h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func main() {
	GetFreakShowWorks()
}
