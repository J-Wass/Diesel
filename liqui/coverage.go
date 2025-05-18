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

	octaneGGLink := utils.Query(liquipediaHTML, "a[href*=\"octane.gg/events\"]")
	octaneGGHref := "#"
	if octaneGGLink != nil {
		octaneGGHref = utils.AttrOr(octaneGGLink, "href", "")
		if octaneGGHref != "" {
			coverageItems = append(coverageItems, fmt.Sprintf("[**Octane.gg**](%s)", octaneGGHref))
		}
	}

	startGGLink := utils.Query(liquipediaHTML, "a[href*=\"start.gg/tournament\"]")
	startGGHref := "#"
	if startGGLink != nil {
		startGGHref = utils.AttrOr(startGGLink, "href", "")
		if startGGHref != "" {
			coverageItems = append(coverageItems, fmt.Sprintf("[**Start.gg**](%s)", startGGHref))
		}
	}

	coverageItems = append(coverageItems, "[**Pickstop.gg**](https://pickstop.gg/rl)")

	blasttvLink := utils.Query(liquipediaHTML, "a[href*=\"blast.tv/rl/tournament\"]")
	blasttvHref := "#"
	if blasttvLink != nil {
		blasttvHref = utils.AttrOr(blasttvLink, "href", "")
		if blasttvHref != "" {
			coverageItems = append(coverageItems, fmt.Sprintf("[**Blast.tv**](%s)", blasttvHref))
		}
	}

	markdown := ""

	markdown += strings.Join(coverageItems, " **/ /** ")
	return markdown
}
