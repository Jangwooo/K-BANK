package bank

import (
	"errors"
	"net/http"

	"K-BANK/model"
	"K-BANK/model/DAO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAccounts is 자신의 계정에 등록된 모든 계좌를 불러오는것
func GetAccounts(c *gin.Context) {
	type account struct {
		BankCode      string `json:"bank_code,omitempty"`
		AccountNumber string `json:"account_number,omitempty"`
		Nickname      string `json:"nickname,omitempty"`
		Balance       int    `json:"balance,omitempty"`
	}
	type response struct {
		Msg      string     `json:"msg"`
		Accounts *[]account `json:"accounts,omitempty"`
	}

	res := new(response)
	res.Accounts = new([]account)

	var accounts []DAO.CheckingAccount
	var otherAccounts []DAO.AnotherAccount

	model.DB.Select("id, bank_id, account_nickname, balance").Find(&accounts, "user_id = ?", c.GetHeader("user_id"))
	model.DB.Select("id, bank_id").Find(&otherAccounts, "user_id = ?", c.GetHeader("user_id"))

	for _, v := range accounts {
		temp := account{
			BankCode:      v.BankID,
			AccountNumber: v.ID,
			Nickname:      v.AccountNickname.String,
			Balance:       v.Balance,
		}
		*res.Accounts = append(*res.Accounts, temp)
	}

	for _, v := range otherAccounts {
		temp := account{
			BankCode:      v.BankID,
			AccountNumber: v.ID,
		}
		*res.Accounts = append(*res.Accounts, temp)
	}

	if len(*res.Accounts) == 0 {
		res.Msg = "계좌가 없습니다"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "성공"
	c.JSON(http.StatusOK, res)
}

// GetAccount is 특정 계좌의 정보를 가져오는것
func GetAccount(c *gin.Context) {
	type response struct {
		Msg     string      `json:"msg"`
		Account interface{} `json:"account,omitempty"`
	}
	res := response{}

	a := new(DAO.CheckingAccount)
	aID := c.Param("account_id")

	if err := model.DB.Take(&a, "id = ?", aID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Msg = "계좌가 없습니다"
			c.JSON(http.StatusOK, res)
			return
		}
		panic(err)
	}

	res.Msg = "성공"
	res.Account = a
	c.JSON(http.StatusOK, res)
}
