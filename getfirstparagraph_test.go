package main

import (
	"testing"
)

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetNoParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			Main paragraph.
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := ""
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}