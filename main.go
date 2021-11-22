package main

import (
	"log"

	"K-BANK/controller"
	"K-BANK/lib"
	"K-BANK/middleware"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	model.Connect()
	lib.CreateCipher()

	r := gin.New()

	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("요청에 실패했습니다 err : %s", err),
		})
	}))

	r.Use(gin.Logger())
	r.Use(cors.Middleware(cors.Config{
		ValidateHeaders: false,
		Origins:         "*",
		RequestHeaders: "Origin, Authorization, Content-Type, Referer, Accept, User-Agent, Accept-Encoding, " +
			"Accept-Language, Cache-Control, Connection, Host, Pragma, Sec-Fetch-Mode, access_token, trade_token",
		ExposedHeaders: "",
		Methods:        "GET, PUT, POST, DELETE",
		MaxAge:         50 * time.Second,
		Credentials:    true,
	}))

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
