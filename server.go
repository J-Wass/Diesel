package main

import (
	"encoding/base64"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	liqui "diesel/liqui"
)

func main() {
	r := gin.Default()

	r.GET("/swiss/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		rootNode, _ := liqui.RootDOMNodeForUrl(string(decodedUrl))
		markdown := liqui.Swiss(rootNode)
		c.String(http.StatusOK, markdown)
	})

	r.GET("/streams/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		rootNode, _ := liqui.RootDOMNodeForUrl(string(decodedUrl))
		markdown := liqui.Streams(rootNode)
		c.String(http.StatusOK, markdown)
	})

	r.GET("/makethread/:url/template/:template", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		rootNode, _ := liqui.RootDOMNodeForUrl(string(decodedUrl))

		base64Template := c.Param("template")
		decodedTemplate, _ := base64.StdEncoding.DecodeString(base64Template)

		markdown := liqui.MakeThread(rootNode, string(decodedTemplate))
		c.String(http.StatusOK, markdown)
	})

	r.GET("/healthcheck", func(c *gin.Context) {
		if liqui.CacheWrites == 0 || liqui.CacheLookups == 0{
			c.JSON(200, gin.H{
				"cache size": len(liqui.DOMCache),
			})
		}else{
			hitRate := float64(liqui.CacheHits)/float64(liqui.CacheLookups)
			thrashRate := float64(liqui.CacheThrashes)/float64(liqui.CacheLookups)
			readWriteRate := float64(liqui.CacheLookups)/float64(liqui.CacheWrites)
			c.JSON(200, gin.H{
				"cache size": len(liqui.DOMCache),
				"cache hit rate": math.Round(hitRate*100)/100,
				"cache thrash rate": math.Round(thrashRate*100)/100,
				"cache read/write rate": math.Round(readWriteRate*100)/100,
			})
		}
	})

	r.GET("/mvp_candidates/:url/teams_allowed/:teams_allowed", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		teamsAllowed, _ := strconv.Atoi(c.Param("teams_allowed"))
		rootNode, _ := liqui.RootDOMNodeForUrl(string(decodedUrl))
		candidates := liqui.MVPCandidates(rootNode, teamsAllowed)
		c.String(http.StatusOK, candidates)
	})

	r.Run()
}

