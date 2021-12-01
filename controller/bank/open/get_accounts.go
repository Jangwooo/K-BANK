package open

import (
	"net/http"

	"K-BANK/model"
	"K-BANK/model/DAO"
	"github.com/gin-gonic/gin"
)

// GetAccounts is 다른 은행에서 이 사람이 가지고 있는 모든 계좌를 조회할때 사용하는 것
func GetAccounts(c *gin.Context) {
	type response struct {
		Msg      string                `json:"msg,omitempty"`
		Accounts []DAO.CheckingAccount `json:"accounts,omitempty"`
	}
	res := response{}

	phoneNumber := c.Param("phone_number")

	var count int64
	subQuery := model.DB.Select("id").Where("phone_number = ?", phoneNumber).Table("users")
	err := model.DB.Where("user_id = (?)", subQuery).Find(&res.Accounts).Count(&count).Error

	if err != nil {
		panic(err)
	}

	if count == 0 {
		res.Msg = "계좌가 없습니다"
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)

}
