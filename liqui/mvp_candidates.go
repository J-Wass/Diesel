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
				// Build list of candidates for the team.
				var candidates []string
				for _, player := range QueryAll(playerTable, "tr"){
					// Skip coaches & subs.
					role := Query(player, "th")
					if role == nil{
						continue
					}
					roleText := role.FirstChild.Data
					if roleText != "1" && roleText != "2" && roleText != "3"{
						continue
					}

					// Add player & teamname to map.
					playerName := Query(player, "td > a").FirstChild.Data
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
	prizepool := Query(doc, "div.general-collapsible.prizepooltable")

	rows := QueryAll(prizepool,  "div.csstable-widget-row span.name")
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

	return StringToBase64(strings.Join(eligibleCandidates, "\n"))
}