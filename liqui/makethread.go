package liqui

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"golang.org/x/net/html"
)

//go:embed game_thread_templates
var gameThreadTemplates embed.FS

func GetTemplates() (map[string]string, []string, error) {
	templates := make(map[string]string)
	templateNames := make([]string, 0)

	// Get the embedded FS for the template directory.
	templateFS, err := fs.Sub(gameThreadTemplates, "game_thread_templates")
	if err != nil {
		return nil, nil, err
	}
	// Iterate all templates in the template dir.
	templateFiles, err := fs.ReadDir(templateFS, ".")
	if err != nil {
		return nil, nil, err
	}
	for _, templateFile := range templateFiles {
		filename := templateFile.Name()
		templateString, err := fs.ReadFile(templateFS, filename)
		if err != nil {
			return nil, nil, err
		}
		templates[filename] = string(templateString)
		templateNames = append(templateNames, filename)
	}

	// Return mapping from template name -> template markdown, and also the list of all templates.
	return templates, templateNames, nil
}

func MakeThread(liquipediaHTML *html.Node, liquiUrl string, templateName string, dayNumber int) string {
	// Get all templates. Todo - just read this at server-start time.
	templates, templateNames, _ := GetTemplates()

	threadMarkdown, ok := templates[templateName]
	if !ok {
		return fmt.Sprintf("Couldn't find a template with name %s. Only found %s.", templateName, strings.Join(templateNames, ", "))
	}

	// Find macros in template and replace with read data.
	if strings.Contains(threadMarkdown, "{TITLE}"){
		titleMarkdown := Title(liquipediaHTML)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{TITLE}", titleMarkdown)
	}

	if strings.Contains(threadMarkdown, "{GROUPS}"){
		groupsMarkdown := Groups(liquipediaHTML, liquiUrl)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{GROUPS}", groupsMarkdown)
	}

	if strings.Contains(threadMarkdown, "{SWISS}"){
		swissMarkdown := Swiss(liquipediaHTML)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{SWISS}", swissMarkdown)
	}

	if strings.Contains(threadMarkdown, "{BRACKET}"){
		bracketMarkdown := Bracket(liquipediaHTML, liquiUrl, dayNumber)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{BRACKET}", bracketMarkdown)
	}

	if strings.Contains(threadMarkdown, "{COVERAGE}"){
		coverageMarkdown := Coverage(liquipediaHTML, liquiUrl)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{COVERAGE}", coverageMarkdown)
	}

	if strings.Contains(threadMarkdown, "{PRIZEPOOL}"){
		prizepoolMarkdown := Prizepool(liquipediaHTML)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{PRIZEPOOL}", prizepoolMarkdown)
	}

	if strings.Contains(threadMarkdown, "{STREAMS}"){
		streamsMarkdown := Streams(liquipediaHTML)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{STREAMS}", streamsMarkdown)
	}

	if strings.Contains(threadMarkdown, "{SCHEDULE}"){
		scheduleMarkdown := Schedule(liquipediaHTML, dayNumber)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{SCHEDULE}", scheduleMarkdown)
	}
	
	if strings.Contains(threadMarkdown, "{BROADCAST}"){
		broadcastMarkdown := Broadcast(liquipediaHTML)
		threadMarkdown = strings.ReplaceAll(threadMarkdown, "{BROADCAST}", broadcastMarkdown)
	}
	
	return threadMarkdown
}