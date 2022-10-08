package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/andybalholm/cascadia"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func main() {
	r := gin.Default()

	r.GET("/swiss/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		markdown := swiss(string(decodedUrl))
		c.String(http.StatusOK, markdown)
	})

	r.Run()
}

// Helper method to insert into the middle of a slice.
func insert(a []string, index int, value string) []string {
  if len(a) == index { // nil or empty slice or after last element
      return append(a, value)
  }
  a = append(a[:index+1], a[index:]...) // index < len(a)
  a[index] = value
  return a
}

func Query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}
	}
	return cascadia.Query(n, sel)
}

func QueryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, sel)
}

func AttrOr(n *html.Node, attrName, or string) string {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}
	return or
}

func swiss(liquipediaUrl string) string {
	res, err := http.Get(liquipediaUrl)
	if err != nil {
		return fmt.Sprintf("Error: Couldn't fetch %s.", liquipediaUrl)
	}
	if res.StatusCode != 200 {
		return fmt.Sprintf("Error: Page returned a bad status code: %d", res.StatusCode)
	}

	b, _ := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf("Error: Couldn't read %s's response.", liquipediaUrl)
	}

	doc, _ := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		return fmt.Sprintf("Error: Couldn't parse %s's response.", liquipediaUrl)
	}

	// The indicator that each cell in the swiss table starts with.
	indicator := map[string]string{
		"swisstable-bgc-win":  "✔️",
		"swisstable-bgc-lose": "❌",
		"swisstable-bgc-":     "",
	}

	// Mapping from team name to team acronym.
	acronymMap := map[string]string{}

	// Iterate the teams and generate the acronymMap
	for _, team := range QueryAll(doc, "div.brkts-matchlist-cell.brkts-matchlist-opponent") {
		teamName := AttrOr(team, "aria-label", "")
		if teamName == "" {
			continue
		}
		match := regexp.MustCompile(`\s[\d]{4}`)
		finalTeamName := strings.ToLower(match.ReplaceAllString(teamName, ""))

		acronym := Query(team, "span.name").FirstChild.Data
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
	for _, swiss_table := range QueryAll(doc, "table.swisstable") {
		var rows []string
		rows = append(rows, "|**#**|**Teams**|**W-L**|**Round 1**|**Round 2**|**Round 3**|**Round 4**|**Round 5**|")
		rows = append(rows, "|:-|:-|:-|:-|:-|:-|:-|:-|")

    // Iterate each row.
		for i, t := range QueryAll(swiss_table, "tr") {
      // First row is just headers.
			if i == 0 {
				continue
			}

			var row []string
			row = append(row, strings.Replace(Query(t, "th").FirstChild.Data, ".", " ", -1))

      // Get the team name and link for each row.
			teamName := nameToAcronym(Query(t, "span.team-template-text").FirstChild.FirstChild.Data)
			teamLink := Query(t, "span.team-template-text a")

			teamMarkdown := ""
			if teamLink == nil {
				teamMarkdown = "**" + teamName + "**"
			} else {
				href := AttrOr(teamLink, "href", "#")
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
			row = append(row, "**"+Query(t, "b").FirstChild.Data+"**")
			for _, td := range QueryAll(t, "td[class^=\"swisstable-bgc\"]") {
        // Create a scoreline, such as "✔️ 3-0 DIG"

        // Get check or x-out indicator.
				match := indicator[AttrOr(td, "class", "#")]

        // Get scoreline.
        spans := QueryAll(td, "span")
        score := spans[len(spans)-1].FirstChild.Data
        if score != "img"{
          match += " " + strings.Replace(score, ":", "-", -1)
        }

        // Get opposing team.
				otherTeam := Query(td, "span[class^=\"team-template\"] a")
				if otherTeam != nil {
					match += " " + nameToAcronym(AttrOr(otherTeam, "title", ""))
				}
				row = append(row, match)
			}
      // Combine all scorelines using reddit markdown.
			rows = append(rows, strings.Join(row, "|"))
		}
    // Combine all rows to make the swiss table.
		rows = insert(rows, int(len(rows)/2)+1, "|-|- - - - -|- - -||||||")
		tables = append(tables, strings.Join(rows, "\n"))
	}

  // Encode into base64 so that the markdown can travel safely back to reddit.
  finalMarkdown := strings.Join(tables, "\n")
  fmt.Println(finalMarkdown)
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(finalMarkdown))
	return base64Encoded
}
