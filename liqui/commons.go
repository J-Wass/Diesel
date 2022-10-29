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

// Contains a mapping from [string concatenation of "Last Modified" & URL] to [the *html.Node pointing to the root html node of that URL]
var DOMCache = make(map[string]*html.Node)
var CacheHits = 0
var CacheLookups = 0
var CacheWrites = 0
var CacheThrashes = 0

func StringToBase64(str string) string{
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Returns an html.Node for the supplied url. Either returns a Node or an error.
func RootDOMNodeForUrl(url string) (*html.Node, error){
	res, err := http.Get(url)

	// Check if date modified + url are in the cache before parsing.
	CacheLookups++
	lastModified := res.Header["Last-Modified"][0]
	cacheKey := fmt.Sprintf("%s%s", lastModified, url)
	if cachedDOMNode, ok := DOMCache[cacheKey]; ok {
		CacheHits++
		return cachedDOMNode, nil
	}

	// Parse http response into an html node.
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

	// If cache is full, just empty and restart build.
	if len(DOMCache) >= 3{
		CacheThrashes++
		DOMCache = make(map[string]*html.Node)
	}

	// Write to cache, return the parsed document.
	CacheWrites++
	DOMCache[cacheKey] = doc


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