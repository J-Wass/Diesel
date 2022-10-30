package main

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	liqui "diesel/liqui"
	utils "diesel/utils"
)

// Sets up all routes for Diesel.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swiss/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(string(decodedUrl))
		markdown :=  liqui.Swiss(rootNode)
		encodedMarkdown :=  utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/streams/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ :=  utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(string(decodedUrl))
		markdown := liqui.Streams(rootNode)
		encodedMarkdown :=  utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/coverage/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ :=  utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(string(decodedUrl))
		markdown :=  liqui.Coverage(rootNode)
		encodedMarkdown :=  utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/makethread/:url/template/:template", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ :=  utils.DecodedFromBase64(base64Url)
		rootNode, _ := utils.RootDOMNodeForUrl(string(decodedUrl))

		base64Template := c.Param("template")
		decodedTemplate, _ :=  utils.DecodedFromBase64(base64Template)

		markdown :=  liqui.MakeThread(rootNode, string(decodedTemplate))
		encodedMarkdown :=  utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})

	r.GET("/healthcheck", func(c *gin.Context) {
		if utils.CacheWrites == 0 || utils.CacheLookups == 0{
			c.JSON(200, gin.H{
				"cache size": len(utils.DOMCache),
			})
		}else{
			hitRate := float64(utils.CacheHits)/float64(utils.CacheLookups)
			thrashRate := float64(utils.CacheThrashes)/float64(utils.CacheLookups)
			readWriteRate := float64(utils.CacheLookups)/float64(utils.CacheWrites)
			c.JSON(200, gin.H{
				"cache size": len(utils.DOMCache),
				"cache hit rate": math.Round(hitRate*100)/100,
				"cache thrash rate": math.Round(thrashRate*100)/100,
				"cache read/write rate": math.Round(readWriteRate*100)/100,
			})
		}
	})

	r.GET("/mvp_candidates/:url/teams_allowed/:teams_allowed", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ :=  utils.DecodedFromBase64(base64Url)
		teamsAllowed, _ := strconv.Atoi(c.Param("teams_allowed"))
		rootNode, _ := utils.RootDOMNodeForUrl(string(decodedUrl))
		markdown :=  liqui.MVPCandidates(rootNode, teamsAllowed)
		encodedMarkdown :=  utils.EncodedBase64(markdown)
		c.String(http.StatusOK, encodedMarkdown)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}

