package liqui

import (
	"fmt"
	"strings"
)

// Returns \n delimited string of all mvp candidates in the liqui url
func MVPCandidates(liquipediaUrl string, teamsAllowed int) string {
	doc, err := RootDOMNodeForUrl(liquipediaUrl)

	if err != nil{
		return err.Error()
	}

	// Mapping from teamName -> list of players
	teamToPlayerMap := map[string][]string{}
	for _, teamBox := range QueryAll(doc, "div[class^=teamcard-columns]"){
		for _, team := range QueryAll(teamBox, "div.template-box"){
			teamName := Query(team, "center > a").FirstChild.Data
			for _, playerTable := range QueryAll(team, "div.teamcard-inner > table"){
				// Skip subs.
				tableHeader := Query(playerTable, "tbody > tr > th > abbr")
				if tableHeader != nil && AttrOr(tableHeader, "title", "") == "Substitute"{
					continue
				}

				// Build list of candidates for the team.
				var candidates []string
				for _, player := range QueryAll(playerTable, "tr:nth-child(-n+3) td > a"){
					playerName := strings.TrimSpace(player.FirstChild.Data)

					// If playerName is an invalid option, just skip.
					var badRows = []string{"DNP","Ranking","Substitutes","Main Roster",}
					shouldSkip := false
					for _, ignoreName := range badRows{
						if strings.Contains(playerName, ignoreName){
							shouldSkip = true
							break
						}
					}
					if shouldSkip || playerName == ""{
						continue
					}

					// Add player & teamname to map.
					candidates = append(candidates, fmt.Sprintf("%s (%s)", playerName, teamName))
				}
				if _, ok := teamToPlayerMap[teamName]; !ok{
					teamToPlayerMap[teamName] = []string{}
				}
				teamToPlayerMap[teamName] = append(teamToPlayerMap[teamName], candidates...)
			}
		}
	}
	fmt.Println(teamToPlayerMap)

	var mvpCandidateStringBuilder strings.Builder

	return StringToBase64(mvpCandidateStringBuilder.String())
}