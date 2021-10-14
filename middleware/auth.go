package middleware

import (
	"context"
	"net/http"

	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	client := c.Request.Header.Get("client")
	accessToken := c.Request.Header.Get("access_token")

	r, err := model.ConnectRedis(model.AccessTokenDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}
	r2, err := model.ConnectRedis(model.MobileTokenDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}

	switch client {
	case "web":
		if id, err := r.Get(context.Background(), accessToken).Result(); err == nil {
			c.Request.Header.Add("user_id", id)
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "세션이 만료되었습니다",
			})
			c.Abort()
			return
		}
	case "mobile":
		if id, err := r2.Get(context.Background(), accessToken).Result(); err == nil {
			c.Request.Header.Add("user_id", id)
			c.Next()
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "잘못된 접근입니다",
			})
			c.Abort()
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "잘못된 접근입니다",
		})
		c.Abort()
		return
	}
}
