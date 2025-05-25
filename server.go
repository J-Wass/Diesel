package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	liqui "diesel/liqui"
	utils "diesel/utils"
)

// Sets up all routes for Diesel.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/schedule/:url/day/:day", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		dayNumber, _ := strconv.Atoi(c.Param("day"))

		markdown := liqui.Schedule(rootNode, dayNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/schedule/:url/date/:date", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		dateNumber, _ := strconv.Atoi(c.Param("date"))

		markdown := liqui.ScheduleWithDate(rootNode, dateNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/swiss/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Swiss(rootNode)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/broadcast/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Broadcast(rootNode)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/bracket/:url/day/:day", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		dayNumber, _ := strconv.Atoi(c.Param("day"))

		markdown := liqui.Bracket(rootNode, decodedUrl, dayNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/bracket/:url/date/:date", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		dateNumber, _ := strconv.Atoi(c.Param("date"))

		markdown := liqui.BracketWithDate(rootNode, decodedUrl, dateNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/streams/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Streams(rootNode)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/prizepool/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Prizepool(rootNode)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/coverage/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Coverage(rootNode, decodedUrl)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/groups/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.Groups(rootNode, decodedUrl)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/makethread/:url/template/:template/day/:day", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)

		base64Template := c.Param("template")
		decodedTemplate, _ := utils.DecodedFromBase64(base64Template)

		dayNumber, _ := strconv.Atoi(c.Param("day"))

		markdown := liqui.MakeThread(rootNode,decodedUrl, decodedTemplate, dayNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/makethread/:url/template/:template/date/:date", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)

		base64Template := c.Param("template")
		decodedTemplate, _ := utils.DecodedFromBase64(base64Template)

		dateNumber, _ := strconv.Atoi(c.Param("date"))

		markdown := liqui.MakeThreadWithDate(rootNode,decodedUrl, decodedTemplate, dateNumber)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/mvp_candidates/:url/teams_allowed/:teams_allowed", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		teamsAllowed, _ := strconv.Atoi(c.Param("teams_allowed"))
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		markdown := liqui.MVPCandidates(rootNode, teamsAllowed)
		encodedMarkdown := utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/title/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(decodedUrl)
		title := liqui.Title(rootNode)
		encodedMarkdown := utils.EncodedBase64(title)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/templates", func(c *gin.Context) {
		_, templateNames, _ := liqui.GetTemplates()
		encodedMarkdown := utils.EncodedBase64(strings.Join(templateNames, ", "))
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/healthcheck", func(c *gin.Context) {
		if utils.CacheWrites == 0 || utils.CacheLookups == 0 {
			c.JSON(200, gin.H{
				"age": fmt.Sprintf("%s commits old", utils.CommitAge()),
				"cache size": len(utils.DOMCache),
			})
		} else {
			hitRate := float64(utils.CacheHits) / float64(utils.CacheLookups)
			thrashRate := float64(utils.CacheThrashes) / float64(utils.CacheLookups)
			readWriteRate := float64(utils.CacheLookups) / float64(utils.CacheWrites)
			c.JSON(200, gin.H{
				"age": fmt.Sprintf("%s commits old", utils.CommitAge()),
				"cache size":            len(utils.DOMCache),
				"cache hit rate":        math.Round(hitRate*100) / 100,
				"cache thrash rate":     math.Round(thrashRate*100) / 100,
				"cache read/write rate": math.Round(readWriteRate*100) / 100,
			})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
