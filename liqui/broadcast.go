package liqui

import (
	"fmt"
	"sort"
	"strings"

	utils "diesel/utils"

	"golang.org/x/net/html"
)

type broadcastStream struct {
	name     string
	link     string
	platform string
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func getBroadcastsFromLiqui(liquipediaHTML *html.Node) []broadcastStream {
	SUPPORTED_STREAMS := []string{"rocketstreetlive", "rocketleague", "rocketleaguemena", "rocketleaguesam", "rocketleageueoce", "rocketleageueapac", "acgl alpha", "therocketrb", "dreamhack", "rizzo", "rloceania", "oceaniarl", "oxg_esports", "rlesports"}
	domains := []string{"twitch", "youtube"}
	domainBroadcasts := map[string][]string{}
	for _, d := range domains {
		domainBroadcasts[d] = []string{}
	}

	center := utils.Query(liquipediaHTML, ".infobox-center")

	links := utils.QueryAll(center, "a")
	for _, link := range links {
		url := utils.AttrOr(link, "href", "#")
		for _, domain := range domains {
			if strings.Contains(url, domain) {
				domainBroadcasts[domain] = append(domainBroadcasts[domain], url)
			}
		}
	}

	broadcastsStreams := make([]broadcastStream, 0)
	for platform, broadcastUrls := range domainBroadcasts {
		for _, broadcastUrl := range broadcastUrls {
			urlSplit := strings.Split(broadcastUrl, "/")
			streamName := urlSplit[len(urlSplit)-1]
			cleanedStreamName := strings.Title(strings.ReplaceAll(streamName, "@", ""))
			cleanedStreamName = strings.Title(strings.ReplaceAll(cleanedStreamName, "_", " "))
			if !contains(SUPPORTED_STREAMS, strings.ToLower(cleanedStreamName)) {
				cleanedStreamName = "Link"
			}
			broadcastsStreams = append(broadcastsStreams, broadcastStream{name: cleanedStreamName, link: broadcastUrl, platform: strings.Title(platform)})
		}
	}
	sort.Slice(broadcastsStreams, func(i, j int) bool {
		s1 := broadcastsStreams[i]
		s2 := broadcastsStreams[j]
		if s1.platform == s2.platform {
			return s1.name < s2.name
		}
		return s1.platform < s2.platform
	})
	return broadcastsStreams
}

func markdownFromBroadcasts(broadcasts []broadcastStream) string {
	var markdownStringBuilder strings.Builder
	markdownStringBuilder.WriteString("|**Platforms**|**Link**|\n|:-|:-|\n")

	for _, broadcast := range broadcasts {
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
