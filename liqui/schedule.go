package liqui

import (
	utils "diesel/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)


func Schedule(liquipediaHTML *html.Node, dayNumber int) string {
	// Get all dates of games.
	dateTimestamps := make([]int, 0)
	dateObjects := utils.QueryAll(liquipediaHTML, "span.timer-object")
	for _, dateObject := range dateObjects {
		newTimestamp := utils.AttrOr(dateObject, "data-timestamp", "0")
		timestampInt, _ := strconv.Atoi(newTimestamp)
		dateTimestamps = append(dateTimestamps, timestampInt)
	}

	sort.Ints(dateTimestamps)
	
	// Group dates of each game into buckets of days.
	dayBuckets := make([][]time.Time, 0)
	prevTimestamp := -1
	for _, timestamp := range dateTimestamps{
		// Start a new day if @ beginning or new timestamp is far in future
		if prevTimestamp == -1 || timestamp - prevTimestamp > 60 * 60 * 2{
			dayBuckets = append(dayBuckets, make([]time.Time, 0))
		}

		// Convert timestamp to datetime, add to correct day bucket.
		datetime := time.Unix(int64(timestamp), 0)
		dayBuckets[len(dayBuckets)-1] = append(dayBuckets[len(dayBuckets)-1], datetime.UTC())
		prevTimestamp = timestamp
	}

	// Build the markdown.
	var markdownStringBuilder strings.Builder
	markdownStringBuilder.WriteString("||**Day**|**UTC**||")
	markdownStringBuilder.WriteString("\n|:-|:-|:-|:-|")
	for i, day := range dayBuckets{
		hour, minute, _ := day[0].Clock()
		dayOfWeek := day[0].Weekday().String()
		dayMarkdown := fmt.Sprintf("\n|Day %d|%s|[**%d:%02d**](https://www.google.com/search?q=%d:%02d+GMT)|", i+1, dayOfWeek, hour, minute, hour, minute)
		if i == dayNumber-1{
			dayMarkdown += "**Today**|"
		}else{
			dayMarkdown += "|"
		}
		markdownStringBuilder.WriteString(dayMarkdown)

	}

	return markdownStringBuilder.String()
}