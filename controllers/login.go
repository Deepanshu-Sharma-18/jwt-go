package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Deepanshu-Sharma-18/jwt-auth/helpers"
	"github.com/Deepanshu-Sharma-18/jwt-auth/initializers"
	"github.com/Deepanshu-Sharma-18/jwt-auth/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	var foundUser models.User

	defer cancel()
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := initializers.Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

	fmt.Println(err)
	fmt.Println(foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
		return
	}

	check := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if check != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect password hash"})
		return
	}

	token, refreshToken, _ := helpers.GenerateToken(foundUser.Email)

	helpers.UpdateToken(token, refreshToken, foundUser.Email, initializers.Collection)
	err = initializers.Collection.FindOne(ctx, bson.M{"email": foundUser.Email}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": foundUser})

}
