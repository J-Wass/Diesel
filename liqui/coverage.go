package liqui

import (
	utils "diesel/utils"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Coverage(liquipediaHTML *html.Node, liquiURL string) string {
	coverageItems := make([]string, 0)
	coverageItems = append(coverageItems, fmt.Sprintf("[**Liquipedia**](%s)", liquiURL))

	octaneGGLink := utils.Query(liquipediaHTML, "a[href^=\"https://octane.gg/events\"]")
	octaneGGHref := "#"
	if octaneGGLink != nil {
		octaneGGHref = utils.AttrOr(octaneGGLink, "href", "")
		if octaneGGHref != "" {
			coverageItems = append(coverageItems, fmt.Sprintf("[**Octane.gg**](%s)", octaneGGHref))
		}
	}

	startGGLink := utils.Query(liquipediaHTML, "a[href^=\"https://www.start.gg/tournament\"]")
	startGGHref := "#"
	if startGGLink != nil {
		startGGHref = utils.AttrOr(startGGLink, "href", "")
		if startGGHref != "" {
			coverageItems = append(coverageItems, fmt.Sprintf("[**Start.gg**](%s)", startGGHref))
		}
	}

	markdown := ""

	markdown += strings.Join(coverageItems, " **/ /** ")
	return markdown
}
