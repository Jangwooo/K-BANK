package main

import (
	"K-BANK/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.POST("/signUp", controller.SignUpHandler)
		}
	}
	r.Static("images", "./images/")

	r.Run()
}
