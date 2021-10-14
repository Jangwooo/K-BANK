package main

import (
	"log"

	"K-BANK/controller"
	"K-BANK/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.POST("/signUp", controller.SignUpHandler)
			userAPI.GET("/identity", controller.Identity)
			userAPI.POST("/login", controller.LoginHandler)
		}

		bankAPI := api.Group("/banking").Use(middleware.Auth)
		{
			bankAPI.POST("/account", controller.OpenAccountHandler)
		}
	}
	r.Static("images", "./images/")

	r.Run()
}
