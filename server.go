package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	liqui "diesel/liqui"
)

func main() {
	r := gin.Default()

	r.GET("/swiss/:url", func(c *gin.Context) {
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		markdown := liqui.Swiss(string(decodedUrl))
		c.String(http.StatusOK, markdown)
	})

	r.GET("/streams/:url", func(c *gin.Context) {
		fmt.Println("loading /streams")
		base64Url := c.Param("url")
		decodedUrl, _ := base64.StdEncoding.DecodeString(base64Url)
		markdown := liqui.Streams(string(decodedUrl))
		c.String(http.StatusOK, markdown)
	})

	r.Run()
}

