package main

import (
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("ERROR -- Issue reading html from site: %v\n", err)
		return ""
	}
	headingOne := doc.Find("h1")
	headingTwo := doc.Find("h2")
	if headingOne.Text() == "" {
		return headingTwo.Text()
	}
	return headingOne.Text()
}

