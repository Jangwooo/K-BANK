package bank

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"K-BANK/lib"
	"K-BANK/model"
	"K-BANK/model/DAO"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type OpenAccountRequest struct {
	AccountNickname string `json:"account_nickname" `
	Pwd             string `json:"pwd" binding:"required"`
	TradeToken      string `header:"trade_token"`
}

// OpenAccountHandler is 일반 입출금 계좌 생성
func OpenAccountHandler(c *gin.Context) {
	type Response struct {
		Msg           string `json:"msg,omitempty"`
		AccountNumber string `json:"account_number,omitempty"`
	}
	res := Response{}

	req := new(OpenAccountRequest)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "리퀘스트 형식이 잘못되었습니다",
		})
		return
	}

	var aid string
	var n int64
	ca := DAO.CheckingAccount{}
	for {
		aid, err = lib.CreateCode()
		if err != nil {
			panic(err)
		}

		err := model.DB.First(&ca, "id = ?", "110"+"513"+aid).Count(&n).Error

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}

		if n == 0 {
			aid = "110" + "513" + aid
			break
		}
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	uid := c.GetHeader("user_id")
	ca = DAO.CheckingAccount{
		ID:       aid,
		BankID:   "110",
		UserID:   uid,
		Password: string(pwd),
		AccountNickname: sql.NullString{
			String: req.AccountNickname,
			Valid:  req.AccountNickname != "",
		},
		Balance:   10000,
		CreatedAt: time.Now(),
		State:     "normal",
		Limit:     10000,
		History: &[]DAO.History{
			{
				TradeType: "New",
				AccountID: uid,
				Client:    "K-BANK",
				CreatedAt: time.Now(),
				Amount:    10000,
				Balance:   10000,
				State:     "success",
				ErrorLog:  nil,
			},
		},
	}

	err = model.DB.Create(&ca).Error
	if err != nil {
		panic(err)
	}

	res.Msg = "성공"
	res.AccountNumber = ca.ID

	c.JSON(http.StatusOK, res)
}
