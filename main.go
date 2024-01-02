package main

import (
	"os"

	"github.com/Deepanshu-Sharma-18/jwt-auth/controllers"
	"github.com/Deepanshu-Sharma-18/jwt-auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()

}
