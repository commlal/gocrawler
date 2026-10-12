package main

import (
	"log"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	urls := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Printf("ERROR -- Issue reading html from site: %v\n", err)
		return urls, err
	}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {  	
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}
		relativeURL, err := url.Parse(href)
		if err != nil {
			log.Printf("ERROR -- Could not parse href %q: %v\n", href, err)
			return
		}
		final := baseURL.ResolveReference(relativeURL)
		urls = append(urls, final.String())
    })
	return urls, nil
}