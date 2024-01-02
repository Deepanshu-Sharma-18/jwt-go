package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Deepanshu-Sharma-18/jwt-auth/helpers"
	"github.com/Deepanshu-Sharma-18/jwt-auth/initializers"
	"github.com/Deepanshu-Sharma-18/jwt-auth/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request/body data",
		})

		return
	}
	collection := initializers.Collection

	result := collection.FindOne(c.Request.Context(), bson.M{"email": body.Email})

	if result.Err() == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error while hashing password",
		})

		return
	}

	user := models.User{Email: body.Email, Password: string(hash), Name: body.Name, CreatedAt: time.Now(), UpdatedAt: time.Now(), Token: "", RefreshToken: "",
		ID: primitive.NewObjectID()}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	token, refreshToken, _ := helpers.GenerateToken(body.Email)
	user.Token = token
	user.RefreshToken = refreshToken

	collection.InsertOne(context.TODO(), user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
