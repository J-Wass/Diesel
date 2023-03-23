package liqui

import (
	utils "diesel/utils"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)


func Schedule(liquipediaHTML *html.Node, dayNumber int) string {
	dayBuckets := utils.DayBuckets(liquipediaHTML)

	// Build the markdown.
	var markdownStringBuilder strings.Builder
	markdownStringBuilder.WriteString("||**Day**|**UTC**||")
	markdownStringBuilder.WriteString("\n|:-|:-|:-|:-|")
	for i, day := range dayBuckets{
		hour, minute, _ := day[0].Clock()
		dayOfWeek := day[0].Weekday().String()
		dayMarkdown := fmt.Sprintf("\n|Day %d|%s|[**%d:%02d**](https://www.google.com/search?q=%d:%02d+GMT)|", i+1, dayOfWeek, hour, minute, hour, minute)
		if i < dayNumber - 1{
			dayMarkdown = fmt.Sprintf("\n|~~Day %d~~|~~%s~~|~~[**%d:%02d**](https://www.google.com/search?q=%d:%02d+GMT)~~|", i+1, dayOfWeek, hour, minute, hour, minute)
		}
		if i == dayNumber-1{
			dayMarkdown += "**Today**|"
		} else{
			dayMarkdown += "|"
		}
		markdownStringBuilder.WriteString(dayMarkdown)

	}

	return markdownStringBuilder.String()
}