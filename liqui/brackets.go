package liqui

import (
	"diesel/utils"
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/bykof/gostradamus"
)

type match struct {
	team1Name     string
	team1Score    string
	team2Name     string
	team2Score    string
	gameStartTime time.Time
	isFinished    bool
}

// Returns a datetime off of the time string from liquipedia.
func datetimeFromLiquiTimestring(timestring string, timezone string) time.Time {
	timezoneTokens := strings.Split(timezone, ":")

	// Ensure the timezone string is 0-padded.
	hour := timezoneTokens[0]
	minute := timezoneTokens[1]
	if len(hour) == 2 {
		hour = strings.ReplaceAll(hour, "-", "-0")
		hour = strings.ReplaceAll(hour, "+", "+0")
	}
	cleanedTimezone := hour + ":" + minute

	// Liqui format example: March 26, 2022 - 13:15
	dateTime, err := gostradamus.Parse(timestring+cleanedTimezone, "MMMM DD, YYYY - HH:mmzz")
	if err != nil {
		return time.Now()
	}
	return time.Time(dateTime)
}

// Returns 'hh:mm UTC' from a datetime"""
func timeOfDayFromDatetime(datetime time.Time) string {
	dt := gostradamus.DateTime(datetime)
	return dt.Format("HH:mm UTC")
}

// Returns a list of matches from a given liquipedia html root node.
func matchesForLiquiURL(liquipediaHTML *html.Node, dayNumber int) []match {
	dayBuckets := utils.DayBuckets(liquipediaHTML)

	// List of each match in the bracket.
	matches := make([]match, 0)
	matchElements := utils.QueryAll(liquipediaHTML, ".brkts-round-center")
	for _, matchElement := range matchElements {
		// Check if match is finished yet.
		timerElement := utils.Query(matchElement, ".timer-object")

		// Get timezone info.
		timezone := utils.AttrOr(utils.Query(timerElement, "abbr"), "data-tz", "")
		timezoneString := utils.Query(timerElement, "abbr").FirstChild.Data

		// Get the timestring, convert to time.Time().
		liquiTimestring := timerElement.FirstChild.Data
		liquiTimestring = strings.ReplaceAll(liquiTimestring, timezoneString, "")
		liquiTimestring = strings.TrimSpace(liquiTimestring)
		gameStartTime := datetimeFromLiquiTimestring(liquiTimestring, timezone)

		// Determine if this game start time is within the requested day.
		if dayNumber > len(dayBuckets) {
			return matches
		}
		allowedDatetimes := dayBuckets[dayNumber-1]
		earliestDatetime := allowedDatetimes[0]
		latestDatetime := allowedDatetimes[len(allowedDatetimes)-1]

		gameStartTimeIsAllowed := false
		if earliestDatetime.Before(latestDatetime) {
			gameStartTimeIsAllowed = !gameStartTime.Before(earliestDatetime) && !gameStartTime.After(latestDatetime)
		} else if earliestDatetime.Equal(latestDatetime) {
			gameStartTimeIsAllowed = gameStartTime.Equal(gameStartTime)
		}

		if !gameStartTimeIsAllowed {
			// This game is outside of the requested day, so ignore it.
			continue
		}

		isFinished := utils.AttrOr(timerElement, "data-finished", "") != ""

		teams := utils.QueryAll(matchElement, ".brkts-opponent-entry")
		team1Name := "TBD"
		team1Score := ""
		team2Name := "TBD"
		team2Score := ""

		for i, team := range teams {
			teamName := "TBD"
			teamNameElement := utils.Query(team, ".name")
			if teamNameElement != nil && teamNameElement.FirstChild != nil {
				teamName = teamNameElement.FirstChild.Data
			}

			// Get team score, if it exists.
			teamScoreElement := utils.Query(team, ".brkts-opponent-score-inner")
			teamScore := "-"
			if teamScoreElement != nil && teamScoreElement.FirstChild != nil {
				teamScore = teamScoreElement.FirstChild.Data
				// If this team has won, the score is wrapped in a b element.
				winningTeamScore := utils.Query(teamScoreElement, "b")
				if winningTeamScore != nil {
					teamScore = winningTeamScore.FirstChild.Data
				}
			}

			if i == 0 {
				team1Name = teamName
				team1Score = teamScore
			} else {
				team2Name = teamName
				team2Score = teamScore
			}
		}

		newMatch := match{
			team1Name:     team1Name,
			team2Name:     team2Name,
			team1Score:    team1Score,
			team2Score:    team2Score,
			gameStartTime: gameStartTime,
			isFinished:    isFinished,
		}
		matches = append(matches, newMatch)
	}

	// Return matches, sorted by game start time.
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].gameStartTime.Before(matches[j].gameStartTime)
	})
	return matches
}

// Returns reddit markdown of the matches as a bracket table.
func markdownForMatches(matches []match, liqui_url string) string {
	// Templates for markdown header and row.
	BRACKET_MARKDOWN_HEADER := "|*ELIMINATION*||[**Liquipedia Bracket**]({LIQUI_URL}#Results)|\n|:-|:-|:-|"
	BRACKET_MARKDOWN_ROW_UNFINISHED := "|{TEAM1}|[**{TIMESTRING}**](https://www.google.com/search?q={TIMESTRING})|{TEAM2}|"
	BRACKET_MARKDOWN_ROW_FINISHED := "|{TEAM1}|**{TEAM1_SCORE} - {TEAM2_SCORE}**|{TEAM2}|"

	var finalMarkdown strings.Builder
	header := strings.ReplaceAll(BRACKET_MARKDOWN_HEADER, "{LIQUI_URL}", liqui_url)
	finalMarkdown.WriteString(header)

	// Each match becomes a row in the markdown table.
	for _, match := range matches {
		team1Name := match.team1Name
		team2Name := match.team2Name
		rowMarkdown := BRACKET_MARKDOWN_ROW_UNFINISHED

		// If match is finished, bold the winning team.
		if match.isFinished {
			rowMarkdown = BRACKET_MARKDOWN_ROW_FINISHED
			if match.team1Score > match.team2Score {
				team1Name = fmt.Sprintf("**%s**", team1Name)
			} else {
				team2Name = fmt.Sprintf("**%s**", team2Name)
			}
		}

		rowMarkdown = strings.ReplaceAll(rowMarkdown, "{TEAM1}", team1Name)
		rowMarkdown = strings.ReplaceAll(rowMarkdown, "{TEAM2}", team2Name)
		rowMarkdown = strings.ReplaceAll(rowMarkdown, "{TIMESTRING}", timeOfDayFromDatetime(match.gameStartTime))
		rowMarkdown = strings.ReplaceAll(rowMarkdown, "{TEAM1_SCORE}", match.team1Score)
		rowMarkdown = strings.ReplaceAll(rowMarkdown, "{TEAM2_SCORE}", match.team2Score)

		finalMarkdown.WriteString("\n" + rowMarkdown)
	}

	return finalMarkdown.String()
}

func Bracket(liquipediaHTML *html.Node, liqui_url string, dayNumber int) string {
	matches := matchesForLiquiURL(liquipediaHTML, dayNumber)
	markdown := markdownForMatches(matches, liqui_url)

	return markdown
}
