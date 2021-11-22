package user

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
	type Response struct {
		Msg           string    `json:"msg,omitempty"`
		AccessToken   string    `json:"access_token,omitempty"`
		Expiration    time.Time `json:"expiration,omitempty"`
		PersonalToken string    `json:"personal_token,omitempty"`
	}
	res := Response{}

	req := new(LoginRequest)
	err := c.Bind(req)
	if err != nil {
		res.Msg = "리퀘스트 형식이 잘못되었습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var u model.User

	var count int64
	model.DB.Where("id = ?", req.ID).Find(&u).Count(&count)
	if count == 0 {
		res.Msg = "일치하는 유저가 없습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	match := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Pwd))
	if match != nil {
		res.Msg = "비밀번호가 일치하지 않습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	accessToken := lib.CreateToken()
	refreshToken := lib.CreateToken()

	model.AccessTokenRedis.Set(context.Background(), accessToken, req.ID, time.Hour*2)
	model.PersonalTokenRedis.Set(context.Background(), refreshToken, req.ID, 0)

	res.Msg = "성공"
	res.AccessToken = accessToken
	res.Expiration = time.Now().Add(time.Hour * 2)
	res.PersonalToken = refreshToken

	c.JSON(http.StatusOK, res)
}
