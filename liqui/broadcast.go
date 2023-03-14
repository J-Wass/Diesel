package liqui

import (
	"fmt"
	"strings"

	utils "diesel/utils"

	"golang.org/x/net/html"
)

type broadcastStream struct {
    name string
    link  string
	platform string
}

func getBroadcastsFromLiqui(liquipediaHTML *html.Node) []broadcastStream {
	domains := []string{"twitch", "youtube"}
	domainBroadcasts := map[string][]string{}
	for _, d := range domains {
		domainBroadcasts[d] = []string{}
	}

	center := utils.Query(liquipediaHTML, ".infobox-center")

	links := utils.QueryAll(center,"a")
	for _, link := range links{
	  url := utils.AttrOr(link, "href", "#")
	  for _, domain := range domains{
		  if strings.Contains(url, domain){
			domainBroadcasts[domain] = append(domainBroadcasts[domain], url)
		  }
	  }
	}

	broadcastsStreams := make([]broadcastStream,0)
	for platform, broadcastUrls := range domainBroadcasts {
		for _, broadcastUrl := range broadcastUrls{
			urlSplit := strings.Split(broadcastUrl, "/")
			streamName := urlSplit[len(urlSplit)-1]
			cleanedStreamName := strings.Title(strings.ReplaceAll(streamName, "@", ""))
			broadcastsStreams = append(broadcastsStreams, broadcastStream{name: cleanedStreamName, link: broadcastUrl, platform: strings.Title(platform)})
		}
	}
	fmt.Print(broadcastsStreams)
	return broadcastsStreams
}

func markdownFromBroadcasts(broadcasts []broadcastStream) string{
	var markdownStringBuilder strings.Builder
	markdownStringBuilder.WriteString("# Streams\n\n|**Platforms**|**Link**|\n|:-|:-|\n")

	for _, broadcast := range broadcasts{
		broadcastRow := fmt.Sprintf("|%s|[**%s**](%s)|\n", broadcast.platform, broadcast.name, broadcast.link)
		markdownStringBuilder.WriteString(broadcastRow)
	}

	return markdownStringBuilder.String()
}

func Broadcast(liquipediaHTML *html.Node) string {	
	broadcasts := getBroadcastsFromLiqui(liquipediaHTML)
	markdown := markdownFromBroadcasts(broadcasts)

	return markdown
}