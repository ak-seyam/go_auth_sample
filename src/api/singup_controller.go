package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/A-Siam/go_auth/src/data/system_errors"
	"github.com/A-Siam/go_auth/src/services/signup"
	"github.com/gin-gonic/gin"
)

func SignUpController(g *gin.RouterGroup) {
	g.POST("/singup", func(c *gin.Context) {
		var singUpRequest signup.SignUpRequest
		if err := c.BindJSON(&singUpRequest); err != nil {
			logger.Printf("error in login parsing %s\n", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error parsing json input",
				"date":    time.Now(),
			})
			return
		}
		signUpResponse, err := signup.SingUp(singUpRequest)
		if err != nil {
			var systemError *system_errors.SystemError
			logger.Printf("error in singing up {%s}\n", err.Error())
			if errors.As(err, &systemError) {
				c.JSON(int(systemError.Code), gin.H{
					"message": systemError.Message,
					"date":    time.Now(),
				})
			} else {
				c.JSON(int(http.StatusInternalServerError), gin.H{
					"message": systemError.Message,
					"date":    time.Now(),
				})
			}
			return
		}
		c.JSON(http.StatusCreated, signUpResponse)
	})
}
