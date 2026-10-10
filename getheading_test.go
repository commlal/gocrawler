package main

import (
	"testing"
)

func TestGetHeadingOneFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1><h2>Sub-Heading</h2></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetHeadingTwoFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h2>Sub-Heading</h2></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Sub-Heading"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetNoHeadingFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h3>Sub-sub-Heading</h3></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := ""
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

