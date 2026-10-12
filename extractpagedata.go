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
	baseURL, _ := url.Parse(pageURL)
	links, _ := getURLsFromHTML(html, baseURL)
	imgURLs, _ := getImagesFromHTML(html, baseURL)
	pageData := PageData{
		URL:		pageURL,
		Heading:	getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks: 	links,
		ImageURLs:		imgURLs,
	}
	return pageData
}