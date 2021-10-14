package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

type IdentityRequest struct {
	Name string `json:"name"`
	Ssn  string `json:"ssn"`
}

func UnmarshalIdentityRequest(data []byte) (IdentityRequest, error) {
	var r IdentityRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func Identity(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	req, err := UnmarshalIdentityRequest(body)
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

	var u User
	_ = db.Get(&u, "select * from user where user_id = ?", c.Request.Header.Get("user_id"))

	if req.Name != u.Name || req.Ssn != u.Ssn {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "입력하신 정보가 일치하지 않습니다",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":         u.Name,
		"ssn":          u.Ssn,
		"phone_number": u.PhoneNumber,
	})

}
