package controller

import (
	"context"
	"net/http"
	"time"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

type IdentityRequest struct {
	Name   string `json:"name"`
	SSN    string `json:"ssn"`
	UserID string `header:"user-id"`
}

func Identity(c *gin.Context) {
	req := new(IdentityRequest)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "리퀘스트 형식이 잘못되었습니다",
		})
		return
	}

	var u model.User
	err = model.DB.First(&u, "id = ?", req.UserID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "입력하신 정보와 일치하는 유저가 없습니다",
		})
		return
	}

	ssn, err := lib.Cipher.Decrypt(u.SSN)
	if err != nil {
		_ = c.Error(err)
	}

	if req.Name != u.Name || req.SSN != ssn {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "입력하신 정보가 일치하지 않습니다",
		})
		return
	}

	tradeToken := lib.CreateToken()
	model.TradeTokenRedis.Set(context.Background(), tradeToken, u.ID, time.Minute*5)

	c.JSON(http.StatusOK, gin.H{
		"name":         u.Name,
		"ssn":          ssn,
		"phone_number": u.PhoneNumber,
		"trade-token":  tradeToken,
	})
}
