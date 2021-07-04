package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkdika/goft/pkg/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "assets")

	router.GET("/ping", pingHandler)
	router.GET("/test", handler.TestHandler)
	router.GET("/healthcheck", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}
