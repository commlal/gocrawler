package main

import (
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("ERROR -- Issue reading html from site: %v\n", err)
		return ""
	}
	main := doc.Find("main")
	paragraph := main.Find("p")
	return paragraph.Text()
}