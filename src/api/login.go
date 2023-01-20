package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func RouteAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.GET("/login", func(c *gin.Context) {
		logger.Println("Hello world!!!")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}
