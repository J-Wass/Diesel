package liqui

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type stream struct {
    name string
    link  string
}

func Streams(liquipediaHTML *html.Node) string {
	var markdownStringBuilder strings.Builder
	markdownStringBuilder.WriteString("|||||\n|:-|:-|:-|:-|")
	formattedStreams := make([]stream, 0)
	streamTables := QueryAll(liquipediaHTML, "table.sortable.wikitable")

	for _, table := range streamTables {
		rows := QueryAll(Query(table, "tbody"), "tr")

		teams := QueryAll(rows[0], "td")
		streams := QueryAll(rows[1], "td")
		

		// Iterate all teams, build a stream for them.
		for i, team := range teams{
			teamname := AttrOr(Query(team, "a"), "title", "")
			rawStreamLink := AttrOr(Query(streams[i], "a"), "title", "https://www.twitch.tv/directory/game/Rocket%20League")
			rawStreamLinkSplit := strings.Split(rawStreamLink, "/")
			streamLink := "https://www.twitch.tv/" + rawStreamLinkSplit[len(rawStreamLinkSplit)-1]
			if len(teamname) >0 && len(streamLink) >0{
				newStream := stream{name: teamname, link: streamLink} 
				formattedStreams = append(formattedStreams, newStream)
			}
		}
		
	}
	moreStreams := len(formattedStreams) > 4
	for moreStreams{
		fourTeams := formattedStreams[:4]
		moreStreams = len(formattedStreams) > 4
		if moreStreams{
			formattedStreams = formattedStreams[4:]
		}
	  
		var rowStringBuilder strings.Builder
		for _, team := range fourTeams{
			if len(team.link)>0 && len(team.name) >0{
				rowMarkdown := fmt.Sprintf("|[%s](%s)", team.name, team.link)
				rowStringBuilder.WriteString(rowMarkdown)
			}else{
				rowStringBuilder.WriteString("|")
			}
			
		}        
		markdownStringBuilder.WriteString("\n" + rowStringBuilder.String())
	}

	return StringToBase64(markdownStringBuilder.String())
}