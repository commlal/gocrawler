package main

import (
	"net/url"
)

type PageData struct {
	URL				string
	Heading			string
	FirstParagraph	string
	OutgoingLinks	[]string
	ImageURLs		[]string
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	links, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		links = nil
	}

	imgURLs, _ := getImagesFromHTML(html, baseURL)
	if err != nil {
		imgURLs = nil
	}

	pageData := PageData{
		URL:			pageURL,
		Heading:		heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks: 	links,
		ImageURLs:		imgURLs,
	}
	return pageData
}