package liqui

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func GetTemplates() (map[string]string, []string){
	templates := make(map[string]string)
	templateNames := make([]string, 0)

	// Get the path to the template directory.
	dieselExe, _ := os.Executable()
	dieselPath := filepath.Dir(dieselExe)
	gameThreadTemplatesPath := filepath.Join(dieselPath, "liqui", "game_thread_templates")
	templateDirectory, _ := os.Open(gameThreadTemplatesPath)
	defer templateDirectory.Close()

	// Iterate all templates in the template dir.
	templateFiles, _ := templateDirectory.ReadDir(-1)
	for _, templateFile := range templateFiles {
		filename := templateFile.Name()
		templatePath := filepath.Join(gameThreadTemplatesPath, filename)
		templateString, _ := ioutil.ReadFile(templatePath)
		templates[filename] = string(templateString)
		templateNames = append(templateNames, filename)
	}

	// Return mapping from template name -> template markdown, and also the list of all templates.
	return templates, templateNames
}

func MakeThread(liquipediaHTML *html.Node, liquiUrl string, templateName string, dayNumber int) string {
	// Get all templates. Todo - just read this at server-start time.
	templates, templateNames := GetTemplates()

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