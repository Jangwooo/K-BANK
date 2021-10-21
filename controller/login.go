package controller

import (
	"context"
	"net/http"
	"time"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	ID  string `json:"id" binding:"required"`
	Pwd string `json:"pwd" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	req := new(LoginRequest)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "리퀘스트 형식이 잘못되었습니다",
		})
		return
	}

	var u model.User

	r := model.DB.Where("id = ?", req.ID).Find(&u).RowsAffected
	if r == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "일치하는 유저가 없습니다",
		})
		return
	}

	match := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Pwd))
	if match != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "비밀번호가 일치하지 않습니다",
		})
		return
	}

	accessToken := lib.CreateToken()
	refreshToken := lib.CreateToken()

	model.AccessTokenRedis.Set(context.Background(), accessToken, req.ID, time.Hour*2)
	model.RefreshTokenRedis.Set(context.Background(), refreshToken, req.ID, time.Hour*24*7)

	c.JSON(200, gin.H{
		"refresh_token": refreshToken,
		"expiration":    time.Now().Add(time.Hour * 2),
		"access_token":  accessToken,
	})
}
