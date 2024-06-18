package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	url = "https://kcfrozenmeals.com/collections/hot-snakz-pie"
)

func main() {
	// Get the page
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Check how many pies there are
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".product__title").Each(func(i int, s *goquery.Selection) {
		productName := s.Text()
		productName = strings.TrimSpace(productName)
		fmt.Printf("%v\n", productName)
	})
}
