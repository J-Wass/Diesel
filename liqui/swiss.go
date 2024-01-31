package liqui

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"

	utils "diesel/utils"
)

func Swiss(liquipediaHTML *html.Node) string {
	// The indicator that each cell in the swiss table starts with.
	indicator := map[string]string{
		"swisstable-bgc-win":  "✔️",
		"swisstable-bgc-lose": "❌",
		"swisstable-bgc-":     "",
	}

	// Mapping from team name to team acronym.
	acronymMap := map[string]string{}

	// Iterate the teams and generate the acronymMap
	for _, team := range utils.QueryAll(liquipediaHTML, "div.brkts-matchlist-cell.brkts-matchlist-opponent") {
		teamName := utils.AttrOr(team, "aria-label", "")
		if teamName == "" {
			continue
		}
		match := regexp.MustCompile(`\s[\d]{4}`)
		finalTeamName := strings.ToLower(match.ReplaceAllString(teamName, ""))

		acronym := utils.Query(team, "span.name").FirstChild.Data
		acronymMap[finalTeamName] = acronym
	}

	//Takes a team name and returns it's acroynm. If the acronym isn't found, returns the full name."""
	nameToAcronym := func(teamName string) string {
		normalized := strings.ToLower(teamName)
		acronym, exists := acronymMap[normalized]

		if !exists {
			return teamName
		}
		return acronym
	}

	// Iterate each swiss table.
	var tables []string
	for _, swiss_table := range utils.QueryAll(liquipediaHTML, "table.swisstable") {
		var rows []string
		rows = append(rows, "|**#**|**Teams**|**W-L**|**Round 1**|**Round 2**|**Round 3**|**Round 4**|**Round 5**|")
		rows = append(rows, "|:-|:-|:-|:-|:-|:-|:-|:-|")

		// Iterate each row.
		for i, t := range utils.QueryAll(swiss_table, "tr") {
			// First row is just headers.
			if i == 0 {
				continue
			}

			var row []string
			row = append(row, strings.Replace(utils.Query(t, "th").FirstChild.Data, ".", " ", -1))

			// Get the team name and link for each row.
			teamName := nameToAcronym(utils.Query(t, "span.team-template-text").FirstChild.FirstChild.Data)
			teamLink := utils.Query(t, "span.team-template-text a")

			teamMarkdown := ""
			if teamLink == nil {
				teamMarkdown = "**" + teamName + "**"
			} else {
				href := utils.AttrOr(teamLink, "href", "#")
				href = strings.Replace(href, "(", "\\(", -1)
				href = strings.Replace(href, ")", "\\)", -1)

				// Deal with relative links on liqui.
				if !strings.Contains(href, "https://liquipedia.net") {
					href = "https://liquipedia.net" + href
				}

				teamMarkdown = "[**" + teamName + "**](" + href + ")"
			}

			// Go through each match (column) for the current team (row).
			row = append(row, teamMarkdown)
			row = append(row, "**"+utils.Query(t, "b").FirstChild.Data+"**")
			for _, td := range utils.QueryAll(t, "td[class^=\"swisstable-bgc\"]") {
				// Create a scoreline, such as "✔️ 3-0 DIG"

				// Get check or x-out indicator.
				match := indicator[utils.AttrOr(td, "class", "#")]

				// Get scoreline.
				spans := utils.QueryAll(td, "span")
				score := spans[len(spans)-1].FirstChild.Data
				if score != "img" && score != "a" {
					match += " " + strings.Replace(score, ":", "-", -1)
				}

				// Get opposing team.
				otherTeam := utils.Query(td, "span[class^=\"team-template\"] a")
				if otherTeam != nil {
					match += " " + nameToAcronym(utils.AttrOr(otherTeam, "title", ""))
				}

				row = append(row, match)
			}
			// Combine all scorelines using reddit markdown.
			rows = append(rows, strings.Join(row, "|"))
		}
		// Combine all rows to make the swiss table.
		rows = utils.Insert(rows, int(len(rows)/2)+1, "|-|- - - - -|- - -||||||")
		tables = append(tables, strings.Join(rows, "\n"))
	}

	finalMarkdown := strings.Join(tables, "\n")
	return finalMarkdown
}