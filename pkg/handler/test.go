package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"name": "Maikel Chandika",
	})
}
