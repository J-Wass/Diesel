package liqui

import (
	utils "diesel/utils"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type teamPrize struct {
	teamName  string
	usePlayerName bool
	prize     string
	points    string
	placement string
}

func teamsForHTML(liquipediaHTML *html.Node) []teamPrize {
	// Find the prizepool and iterate each row.
	prizepool := utils.Query(liquipediaHTML, "div.prizepooltable")
	prizepoolRows := utils.QueryAll(prizepool, "div.csstable-widget-row")
	prizes := make([]teamPrize, 0)

	// The order of the prizepool table headers can change.
	// Initialize the headers as -1, if we find them in the headers table we will update.
	headers := utils.QueryAll(liquipediaHTML, "div.prizepooltable-header div.csstable-widget-cell")
	//placeIndex := -1
	dollarPrizeIndex := -1
	rlcsPointsIndex := -1
	//teamNameIndex := -1
	for index, header := range headers {
		column := ""
		if header.FirstChild.Data == "span" {
			column = header.LastChild.FirstChild.FirstChild.Data
		} else {
			column = header.FirstChild.Data
		}
		column = strings.ReplaceAll(html.UnescapeString(column), "\u00A0", "")

		if column == "$" {
			dollarPrizeIndex = index
		} else if column == "RLCS Points" {
			rlcsPointsIndex = index
		}
		fmt.Printf("col %s %s %d %t\n", column, utils.EncodedBase64(column), index, column == "$")
		// } else if column == "Place"{
		// 	placeIndex = index
		// }  else if column == "Participant"{
		// 	teamNameIndex = index
		// }
	}

	for _, row := range prizepoolRows {

		// Skips prizepool rows that are the header or expand menu.
		className := utils.AttrOr(row, "class", "")
		if strings.Contains(className, "prizepooltable-header") || strings.Contains(className, "ppt-toggle-expand") {
			continue
		}

		// Get placement, name, prize, and points for each row. Remove the non-breaking space.

		// The placement (1st, 3rd, 8th, etc) is either under div.prizepooltable-place, or under div.prizepooltable-place > span (the span adds decoration to the placement)
		placement_element := utils.Query(row, "div.prizepooltable-place")
		placement := placement_element.FirstChild.Data
		if utils.Query(placement_element, "span") != nil {
			placement = placement_element.LastChild.Data
		}
		placement = strings.ReplaceAll(placement, "\u00A0", "")

		// Fetch all cells, and rely on indices to get prize & points.
		allCells := utils.QueryAll(row, "div.csstable-widget-cell")
		prize := "N/A"
		if dollarPrizeIndex >= 0 {
			prize = allCells[dollarPrizeIndex].FirstChild.Data

		}

		// Not all tourneys have rlcs points.
		points := "N/A"
		if rlcsPointsIndex >= 0 {
			// points may have an inner div for decoration, so we check if there is an additional child element.
			point_subelement := allCells[rlcsPointsIndex].FirstChild
			points = point_subelement.Data
			if point_subelement.FirstChild != nil {
				points = point_subelement.FirstChild.Data
			}
		}

		// Since some teams are in the same row (and share a prize), we'll have to check them all.
		teams := utils.QueryAll(row, "div.block-team")
		usePlayerName := false // indicates if these are teams or players (1v1)
		if len(teams) == 0{
			// if teams are none, check for players (satisfies 1v1s)
			teams = utils.QueryAll(row, "div.block-player")
			usePlayerName = true
		}
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
				usePlayerName: usePlayerName,
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

	includePoints := teamPrizes[0].points != "N/A"

	PRIZEPOOL_HEADER := "|**Place**|**Prize**|**Team**|\n|:-|:-|:-|"
	PRIZEPOOL_HEADER_PLAYER := "|**Place**|**Prize**|**Player**|\n|:-|:-|:-|"
	PRIZEPOOL_ROW := `|**{PLACEMENT}**|{PRIZE}|{TEAM_NAME}|`
	if includePoints {
		PRIZEPOOL_HEADER = "|**Place**|**Prize**|**Team**|**RLCS Points**|\n|:-|:-|:-|:-|"
		PRIZEPOOL_ROW = `|**{PLACEMENT}**|{PRIZE}|{TEAM_NAME}|{POINTS}|`
	}

	var finalMarkdown strings.Builder
	if teamPrizes[0].usePlayerName{
		finalMarkdown.WriteString(PRIZEPOOL_HEADER_PLAYER)
	}else{
		finalMarkdown.WriteString(PRIZEPOOL_HEADER)
	}

	for _, prize := range teamPrizes {
		newRow := strings.ReplaceAll(PRIZEPOOL_ROW, "{PLACEMENT}", prize.placement)
		newRow = strings.ReplaceAll(newRow, "{PRIZE}", prize.prize)
		newRow = strings.ReplaceAll(newRow, "{TEAM_NAME}", prize.teamName)
		if includePoints {
			newRow = strings.ReplaceAll(newRow, "{POINTS}", prize.points)
		}
		finalMarkdown.WriteString("\n" + newRow)
	}
	return finalMarkdown.String()
}

func Prizepool(liquipediaHTML *html.Node) string {
	teamPrizes := teamsForHTML(liquipediaHTML)
	if len(teamPrizes) == 0 {
		return ""
	}
	return markdownForTeamPrizes(teamPrizes)
}
