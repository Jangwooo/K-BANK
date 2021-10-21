package main

import (
	"log"

	"K-BANK/controller"
	"K-BANK/lib"
	"K-BANK/middleware"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	model.Connect()
	lib.CreateCipher()

	r := gin.Default()

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.POST("/signUp", controller.SignUpHandler)
			userAPI.POST("/login", controller.LoginHandler)
		}

		authAPI := api.Group("/auth").Use(middleware.Auth)
		{
			authAPI.POST("/identity", controller.Identity)
		}

		bankAPI := api.Group("/banking").Use(middleware.Auth)
		{
			bankAPI.POST("/account", controller.OpenAccountHandler)
		}
	}
	r.Static("images", "./images/")

	_ = r.Run(":8000")
}
