package main

import (
	"testing"
	"net/url"
	"reflect"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLAbsoluteThree(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body>
	<a href="https://crawler-test.com/one.html"><span>Boot.dev</span></a>
	<a href="https://crawler-test.com/two.html"><span>Boot.dev</span></a>
	<a href="https://crawler-test.com/three.html"><span>Boot.dev</span></a>
	</body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/one.html", "https://crawler-test.com/two.html", "https://crawler-test.com/three.html"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="/blog/post.html"><span>New Boot.dev Blog Post</span></a></body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/blog/post.html"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelativeThree(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body>
	<a href="/one.html"><span>Boot.dev</span></a>
	<a href="/two.html"><span>Boot.dev</span></a>
	<a href="/three.html"><span>Boot.dev</span></a>
	</body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/one.html", "https://crawler-test.com/two.html", "https://crawler-test.com/three.html"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}