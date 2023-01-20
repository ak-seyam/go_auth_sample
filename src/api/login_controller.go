package api

import (
	"net/http"
	"time"

	"github.com/A-Siam/go_auth/src/services/login"
	"github.com/gin-gonic/gin"
)

func LoginController(g *gin.RouterGroup) {
	g.POST("/login", func(c *gin.Context) {
		var loginRequest login.LoginRequest
		if err := c.BindJSON(&loginRequest); err != nil {
			logger.Printf("error in login parsing %s\n", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error parsing json input",
				"date":    time.Now(),
			})
			return
		}
		logger.Printf(
			"login into system with username {%s} and password **** ...\n",
			loginRequest.Username,
		)
		loginResponse, err := login.Login(loginRequest.Username, loginRequest.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, login.NewLoginError(err.Error()))
		}
		c.JSON(http.StatusOK, loginResponse)
	})
}
