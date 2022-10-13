package liqui

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func StringToBase64(str string) string{
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Returns an html.Node for the supplied url. Either returns a Node or an error.
func RootDOMNodeForUrl(url string) (*html.Node, error){
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch %s", url)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s returned a bad status code: %d",url, res.StatusCode)
	}

	b, _ := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s's response", url)
	}

	doc, _ := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s's response", url)
	}

	return doc, nil
}

// Helper method to insert into the middle of a slice.
func Insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func Query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}
	}
	return cascadia.Query(n, sel)
}

func QueryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, sel)
}

func AttrOr(n *html.Node, attrName, or string) string {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}
	return or
}