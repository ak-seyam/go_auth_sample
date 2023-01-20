package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func RouteAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	LoginController(auth)
	SignUpController(auth)
}
