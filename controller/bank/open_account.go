package bank

import (
	"database/sql"
	"net/http"
	"time"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type OpenAccountRequest struct {
	AccountNickname string `json:"account_nickname" `
	Pwd             string `json:"pwd" binding:"required"`
	TradeToken      string `header:"trade_token"`
}

// OpenAccountHandler : 일반 입출금 계좌 생성
func OpenAccountHandler(c *gin.Context) {
	req := new(OpenAccountRequest)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "리퀘스트 형식이 잘못되었습니다",
		})
		return
	}

	var aNum string
	var count int64
	ca := model.CheckingAccount{}
	for {
		aNum, err = lib.CreateCode()
		if err != nil {
			panic(err)
		}

		model.DB.First(&ca, "id = ?", "110513"+aNum).Count(&count)
		if count == 0 {
			aNum = "110" + "513" + aNum
			break
		}
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	uid := c.GetHeader("user_id")
	ca = model.CheckingAccount{
		ID:       aNum,
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
		History: []model.History{
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

	c.JSON(http.StatusOK, gin.H{
		"msg":     "성공!",
		"account": ca,
	})
}
