package liqui

import (
	utils "diesel/utils"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func Groups(liquipediaHTML *html.Node, liquiURL string) string {

	type team struct {
		teamName string
		teamLink  string
		matchRecord string
		plusMinus string
	}

	type group struct{
		groupName string
		teams []team
	}

	groups := make([]group, 0)

	// Iterate each group table.
	tables := utils.QueryAll(liquipediaHTML, "table.grouptable")
	for _, table := range tables{
		teams := make([]team, 0)
		
		// Get the group name and start iterating each row of the group.
		groupName := utils.Query(table,"tr:nth-child(1) th span").FirstChild.Data
		groupRows := utils.QueryAll(table, "tr:nth-child(n+2)")
		for _, row := range groupRows{
			// Fetch data about each team by scanning the rows of the group table.
			rowData := utils.QueryAll(row, "td")

			teamName := utils.Query(rowData[0], "span.team-template-text").FirstChild.FirstChild.Data
			teamName = strings.TrimSpace(teamName)

			teamLinkElement := utils.Query(rowData[0], ".team-template-text a")
			teamLink := "https://liquipedia.net" + utils.AttrOr(teamLinkElement, "href", "#")

			
			matchRecord := rowData[1].FirstChild.FirstChild.Data
			plusMinus := rowData[3].FirstChild.FirstChild.Data

			newTeam := team{teamName: teamName, teamLink: teamLink, matchRecord: matchRecord, plusMinus: plusMinus}
			teams = append(teams, newTeam)
		}

		// Once all teams are gathered, create a new group and add to list of groups.
		newGroup := group{groupName: groupName, teams: teams}
		groups = append(groups, newGroup)
	}

	GROUP_TEMPLATE_HEADER := "|||||\n|:-|:-|:-|:-|\n|**#**|**{GROUP_NAME}** &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; &#x200B; |**Matches** |**Game Diff** |"
    GROUP_TEMPLATE_ROW := "|{PLACEMENT}|[**{NAME}**]({LINK})|{MATCH_RECORD}|{PLUS_MINUS}|"

	var finalMarkdown strings.Builder
	for _, group := range groups{
		// Create the table header.
		var groupMarkdown strings.Builder
		header := strings.ReplaceAll(GROUP_TEMPLATE_HEADER, "{GROUP_NAME}", group.groupName)
		groupMarkdown.WriteString(header)

		// Create a markdown row for each team in the group.
		for i, team := range group.teams{
			teamRow := GROUP_TEMPLATE_ROW
			teamRow = strings.ReplaceAll(teamRow, "{PLACEMENT}", strconv.Itoa(i+1))
			teamRow = strings.ReplaceAll(teamRow, "{NAME}", team.teamName)
			teamRow = strings.ReplaceAll(teamRow, "{LINK}", team.teamLink)
			teamRow = strings.ReplaceAll(teamRow, "{MATCH_RECORD}", team.matchRecord)
			teamRow = strings.ReplaceAll(teamRow, "{PLUS_MINUS}", team.plusMinus)

			groupMarkdown.WriteString("\n" + teamRow)
		}
		finalMarkdown.WriteString(groupMarkdown.String() + "\n\n&#x200B;\n\n")
	}

	return finalMarkdown.String()
}
