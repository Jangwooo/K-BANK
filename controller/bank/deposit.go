package bank

import (
	"errors"
	"net/http"

	"K-BANK/model"
	"K-BANK/model/DAO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Deposit(c *gin.Context) {
	type request struct {
		Sender struct {
			BankID    string `json:"bank_id"`
			AccountID string `json:"account_id"`
			Name      string `json:"name"`
		} `json:"sender"`
		Receiver struct {
			AccountID string `json:"account_id"`
		} `json:"receiver"`
		Amount int `json:"amount"`
	}

	type response struct {
		Msg string `json:"msg"`
	}

	req := new(request)
	_ = c.Bind(req)

	res := new(response)

	a := DAO.CheckingAccount{}
	if err := model.DB.Find(&a, "id = ?", req.Receiver.AccountID).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		res.Msg = "일치하는 계좌가 없습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := model.DB.Model(&a).Update("balance", a.Balance+req.Amount).Error; err != nil {
		panic(err)
	}
	if err := model.DB.Create(&DAO.History{
		TradeType: "deposit",
		AccountID: req.Receiver.AccountID,
		Client:    req.Sender.Name,
		Amount:    req.Amount,
		Balance:   a.Balance + req.Amount,
		State:     "success",
		ErrorLog:  nil,
	}); err != nil {
		panic(err)
	}

	res.Msg = "성공!"
	c.JSON(http.StatusOK, res)

}
