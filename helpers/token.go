package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Deepanshu-Sharma-18/jwt-auth/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(email string) (signedToken string, signedRefreshToken string, err error) {
	claims := &models.SignedDetails{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 24 * 30)},
		},
	}

	refreshClaims := &models.SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 24 * 30)},
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *models.SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt.Time.Compare(time.Now()) > 0 {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func UpdateToken(signedToken string, signedRefreshToken string, email string, collection *mongo.Collection) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"email": email}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := collection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	return
}
