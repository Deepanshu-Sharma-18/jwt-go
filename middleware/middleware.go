package middleware

import (
	"fmt"
	"net/http"

	"github.com/Deepanshu-Sharma-18/jwt-auth/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)

	c.Next()
}
