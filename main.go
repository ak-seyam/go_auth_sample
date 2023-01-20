package main

import (
	"github.com/A-Siam/go_auth/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RouteAuth(router)
	router.Run(":8899")
}
