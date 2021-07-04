package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "assets")

	r.GET("/ping", pingHandler)
	r.GET("/test", testHandler)

	r.Run()
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}

func testHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"name": "Maikel Chandika",
	})
}
