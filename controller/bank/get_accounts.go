package bank

import (
	"net/http"

	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

// GetAccounts is 자신의 계정에 등록된 모든 계좌를 불러오는것
func GetAccounts(c *gin.Context) {
	type account struct {
		BankCode      string `json:"bank_code,omitempty"`
		AccountNumber string `json:"account_number,omitempty"`
	}
	type response struct {
		Msg      string     `json:"msg"`
		Accounts *[]account `json:"accounts,omitempty"`
	}

	res := new(response)
	res.Accounts = new([]account)

	var accounts []model.CheckingAccount
	var otherAccounts []model.AnotherAccount

	model.DB.Select("id, bank_id").Find(&accounts, "user_id = ?", c.GetHeader("user_id"))
	model.DB.Select("id, bank_id").Find(&otherAccounts, "user_id = ?", c.GetHeader("user_id"))

	for _, v := range accounts {
		temp := account{
			BankCode:      v.BankID,
			AccountNumber: v.ID,
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
