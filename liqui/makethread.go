package liqui

import (
	"golang.org/x/net/html"
)


func MakeThread(liquipediaHTML *html.Node, templateMarkdown string) string {
	// TODO implement
	return StringToBase64(templateMarkdown)
}