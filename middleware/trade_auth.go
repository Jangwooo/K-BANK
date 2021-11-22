package middleware

import (
	"context"
	"net/http"

	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

func TradeAuth(c *gin.Context) {
	token := c.GetHeader("trade_token")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "거래 토큰을 넣어 주십시오",
		})
		c.Abort()
		return
	}

	token, err := model.TradeTokenRedis.Get(context.Background(), token).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "올바른 거래토큰이 아니거나 만료된 토큰입니다",
		})
		c.Abort()
		return
	}
	c.Next()
}
