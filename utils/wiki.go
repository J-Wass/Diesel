package utils

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/cascadia"
	"github.com/bykof/gostradamus"
	"golang.org/x/net/html"
)

// Contains a mapping from [string concatenation of "Last Modified" & URL] to [the *html.Node pointing to the root html node of that URL]
var DOMCache = make(map[string]*html.Node)
var CacheHits = 0
var CacheLookups = 0
var CacheWrites = 0
var CacheThrashes = 0

// Returns an html.Node for the supplied url. Either returns a Node or an error.
func RootDOMNodeForUrl(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	CacheLookups++

	// Get the last modified header. If not avail, use time.now()
	lastModifiedHeaders := res.Header["Last-Modified"]
	lastModified := gostradamus.DateTime(time.Now()).Format("ddd, DD MMM YYYY HH:mm:ss ZZZ") //Thu, 02 Feb 2023 14:21:32 GMT
	if lastModifiedHeaders != nil && len(lastModifiedHeaders) > 0 {
		lastModified = lastModifiedHeaders[0]
	}

	// Check if date modified + url are in the cache before parsing.
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
		return nil, fmt.Errorf("%s returned a bad status code: %d", url, res.StatusCode)
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
	if len(DOMCache) >= 3 {
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

// Returns a 2d slice of time.Times. Index i is a list of all event times for day i. Starts at day 0.
func DayBuckets(liquipediaHTML *html.Node) [][]time.Time {
	// Get all dates of games.
	dateTimestamps := make([]int, 0)
	dateObjects := QueryAll(liquipediaHTML, "span.timer-object")
	for _, dateObject := range dateObjects {
		newTimestamp := AttrOr(dateObject, "data-timestamp", "0")
		timestampInt, _ := strconv.Atoi(newTimestamp)
		dateTimestamps = append(dateTimestamps, timestampInt)
	}

	sort.Ints(dateTimestamps)

	// Group dates of each game into buckets of days.
	dayBuckets := make([][]time.Time, 0)
	prevTimestamp := -1
	for _, timestamp := range dateTimestamps {
		// Start a new day if @ beginning or new timestamp is far in future
		if prevTimestamp == -1 || timestamp-prevTimestamp > 60*60*2 {
			dayBuckets = append(dayBuckets, make([]time.Time, 0))
		}

		// Convert timestamp to datetime, add to correct day bucket.
		datetime := time.Unix(int64(timestamp), 0)
		dayBuckets[len(dayBuckets)-1] = append(dayBuckets[len(dayBuckets)-1], datetime.UTC())
		prevTimestamp = timestamp
	}

	return dayBuckets
}
