package liqui

import (
	utils "diesel/utils"
	"strings"

	"golang.org/x/net/html"
)

type teamPrize struct {
	teamName  string
	prize     string
	points    string
	placement string
}

func teamsForHTML(liquipediaHTML *html.Node) []teamPrize {
	// Find the prizepool and iterate each row.
	prizepool := utils.Query(liquipediaHTML, "div.prizepooltable")
	prizepoolRows := utils.QueryAll(prizepool, "div.csstable-widget-row")
	prizes := make([]teamPrize, 0)
	for _, row := range prizepoolRows {

		// Skips prizepool rows that are the header or expand menu.
		className := utils.AttrOr(row, "class", "")
		if strings.Contains(className, "prizepooltable-header") || strings.Contains(className, "ppt-toggle-expand") {
			continue
		}

		// Get placement, name, prize, and points for each row. Remove the non-breaking space.
		placement := utils.Query(row, "div.prizepooltable-place > div").LastChild.Data
		placement = strings.ReplaceAll(placement, "\u00A0", "")

		// Fetch all cells, and rely on indices to get prize & points.
		allCells := utils.QueryAll(row, "div.csstable-widget-cell")
		prize := allCells[1].FirstChild.Data
		points := allCells[2].FirstChild.FirstChild.Data

		// Since some teams are in the same row (and share a prize), we'll have to check them all.
		teams := utils.QueryAll(row, "div.block-team")
		for _, team := range teams {
			teamName := "TBD"
			teamNameElement := utils.Query(team, "span.name > a")
			if teamNameElement != nil {
				teamName = teamNameElement.FirstChild.Data
			}
			newPrize := teamPrize{
				teamName:  teamName,
				prize:     prize,
				placement: placement,
				points:    points,
			}
			prizes = append(prizes, newPrize)
		}

		// Only get 8 prizepool lines.
		if len(prizes) >= 8 {
			break
		}
	}
	return prizes
}

func markdownForTeamPrizes(teamPrizes []teamPrize) string {

	PRIZEPOOL_HEADER := "|**Place**|**Prize**|**Team**|**RLCS Points**|\n|:-|:-|:-|:-|"
	PRIZEPOOL_ROW := `|**{PLACEMENT}**|{PRIZE}|{TEAM_NAME}|+{POINTS} **()**|`

	var finalMarkdown strings.Builder
	finalMarkdown.WriteString(PRIZEPOOL_HEADER)
	for _, prize := range teamPrizes {
		newRow := strings.ReplaceAll(PRIZEPOOL_ROW, "{PLACEMENT}", prize.placement)
		newRow = strings.ReplaceAll(newRow, "{PRIZE}", prize.prize)
		newRow = strings.ReplaceAll(newRow, "{TEAM_NAME}", prize.teamName)
		newRow = strings.ReplaceAll(newRow, "{POINTS}", prize.points)
		finalMarkdown.WriteString("\n" + newRow)
	}
	return finalMarkdown.String()
}

func Prizepool(liquipediaHTML *html.Node) string {
	teamPrizes := teamsForHTML(liquipediaHTML)
	prizepoolMarkdown := markdownForTeamPrizes(teamPrizes)
	return prizepoolMarkdown
}
