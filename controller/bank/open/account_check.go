package open

import (
	"errors"
	"net/http"

	"K-BANK/model"
	"K-BANK/model/DAO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AccountCheck is 보내는 계좌가 맞는지 확인하는것
func AccountCheck(c *gin.Context) {
	aid := c.Query("account_id")

	var n int64
	type response struct {
		Msg       string `json:"msg,omitempty"`
		Name      string `json:"name,omitempty"`
		AccountID string `json:"account_id,omitempty"`
	}

	res := response{}

	err := model.DB.Model(&DAO.User{}).Select("checking_accounts.id as account_id, name").
		Joins("left join checking_accounts on checking_accounts.user_id = users.id").
		Where("checking_accounts.id = ?", aid).Scan(&res).Count(&n).Error

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		panic(err)
	}

	if n == 0 {
		res.Msg = "해당하는 계좌가 없습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Msg = "성공!"
	c.JSON(http.StatusOK, res)
}
