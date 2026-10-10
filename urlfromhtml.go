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
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {  	
		link, _ := s.Attr("href")
		if !strings.HasPrefix(link, "http") {
			path, _:= url.JoinPath(baseURL.String(), link)
			urls = append(urls, path)
		} else {
			urls = append(urls, link)
		}
    })
	return urls, nil
}