package liqui

import (
	utils "diesel/utils"

	"golang.org/x/net/html"
)

func Title(liquipediaHTML *html.Node) string {

	header := utils.Query(liquipediaHTML, "div.infobox-rocket div.infobox-header")
	return header.LastChild.Data
}
