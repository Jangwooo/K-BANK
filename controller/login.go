package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	ID  string `json:"id"`
	Pwd string `json:"pwd"`
}

func UnmarshalLoginRequest(data []byte) (LoginRequest, error) {
	var r LoginRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func LoginHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	req, err := UnmarshalLoginRequest(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	db, err := model.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	defer db.Close()

	var pwd string
	err = db.Get(&pwd, "SELECT password FROM user where user_id = ?", req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "일치하는 유저가 없습니다",
		})
		return
	}
	match := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(req.Pwd))
	if match != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "비밀번호가 일치하지 않습니다",
		})
		return
	}

	accessToken := lib.CreateToken()
	refreshToken := lib.CreateToken()

	r, err := model.ConnectRedis(model.AccessTokenDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"여기냐?": "그렇다!",
			"msg":  err.Error(),
		})
		return
	}
	r2, err := model.ConnectRedis(model.RefreshTokenDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"여기냐?": "그렇다!",
			"msg":  err.Error(),
		})
		return
	}
	r3, err := model.ConnectRedis(model.MobileTokenDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"여기냐?": "그렇다!",
			"msg":  err.Error(),
		})
		return
	}

	if client := c.Request.Header.Get("client"); client == "mobile" {
		r3.Set(context.Background(), accessToken, req.ID, 0)
		refreshToken = ""
	} else if client == "web" {
		r.Set(context.Background(), accessToken, req.ID, time.Minute*2)
		r2.Set(context.Background(), lib.CreateToken(), req.ID, time.Hour*24*7)
	}

	c.JSON(200, gin.H{
		"refresh_token": refreshToken,
		"expiration":    time.Now().Add(time.Hour * 2),
		"access_token":  accessToken,
	})
}
