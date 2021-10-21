package middleware

import (
	"context"
	"net/http"

	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	accessToken := c.GetHeader("access_token")

	if id, err := model.AccessTokenRedis.Get(context.Background(), accessToken).Result(); err == nil {
		c.Request.Header.Add("user-id", id)
		c.Next()
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "세션이 만료되었습니다",
		})
		c.Abort()
		return
	}
}
