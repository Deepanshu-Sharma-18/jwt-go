package controllers

import (
	"net/http"

	"github.com/Deepanshu-Sharma-18/jwt-auth/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	var body struct {
		Token string `json:"token"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request/body data",
		})

		return
	}

	_, err := helpers.ValidateToken(body.Token)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
	})
}
