package liqui

import (
	"fmt"
	"strings"

	utils "diesel/utils"

	"golang.org/x/net/html"
)

// Returns \n delimited string of all mvp candidates in the liqui url
func MVPCandidates(liquipediaHTML *html.Node, teamsAllowed int) string {
	// Mapping from teamName -> list of players
	teamToPlayerMap := map[string][]string{}
	for _, teamBox := range utils.QueryAll(liquipediaHTML, "div[class^=teamcard-columns]"){
		for _, team := range utils.QueryAll(teamBox, "div.template-box"){
			teamName := utils.Query(team, "center > a").FirstChild.Data
			for _, playerTable := range utils.QueryAll(team, "div.teamcard-inner > table"){
				// Build list of candidates for the team.
				var candidates []string
				for _, player := range utils.QueryAll(playerTable, "tr"){
					// Skip coaches & subs.
					role := utils.Query(player, "th")
					if role == nil{
						continue
					}
					roleText := role.FirstChild.Data
					if roleText != "1" && roleText != "2" && roleText != "3"{
						continue
					}

					// Add player & teamname to map.
					playerName := utils.Query(player, "td > a").FirstChild.Data
					candidates = append(candidates, fmt.Sprintf("%s (%s)", playerName, teamName))
				}
				if _, ok := teamToPlayerMap[teamName]; !ok{
					teamToPlayerMap[teamName] = []string{}
				}
				teamToPlayerMap[teamName] = append(teamToPlayerMap[teamName], candidates...)
			}
		}
	}

	// Iterate prizepool, drawing in players until we cap out at `teamsAllowed`.
	var eligibleCandidates []string
	prizepool := utils.Query(liquipediaHTML, "div.general-collapsible.prizepooltable")

	rows := utils.QueryAll(prizepool,  "div.csstable-widget-row span.name")
	for i, prizepoolRow := range rows{
		if i >= teamsAllowed{
			break
		}
		teamName := strings.TrimSpace(prizepoolRow.FirstChild.FirstChild.Data)
		if teamName == "TBD" || teamName == ""{
			continue
		}
		eligibleCandidates = append(eligibleCandidates, teamToPlayerMap[teamName]...)
	}

	return strings.Join(eligibleCandidates, "\n")
}