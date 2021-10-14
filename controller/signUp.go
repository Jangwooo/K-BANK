package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string `json:"id" db:"user_id"`
	Pwd         string `json:"pwd" db:"password"`
	SimplePwd   string `json:"simple_pwd" db:"phone_number"`
	PhoneNumber string `json:"phone_number" db:"simple_pw"`
	Ssn         string `json:"ssn" db:"SSN"`
	Name        string `json:"name" db:"name"`
	Nickname    string `json:"nickname" db:"nickname"`
	Agree       string `json:"agree" db:"agree"`
	UserType    string `json:"user_type" db:"user_type"`
}

func UnmarshalSignInRequest(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func SignUpHandler(c *gin.Context) {
	db, err := model.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "서버에서! 에러가! 났다!",
		})
		return
	}
	defer db.Close()

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "서버에서! 에러가! 났다!",
		})
		return
	}

	req, err := UnmarshalSignInRequest(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "서버에서! 에러가! 났다!",
		})
		return
	}

	result := lib.DuplicateCheck("user", "user_id", req.ID)
	if result == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ID 중복됨",
		})
		return
	}

	result = lib.DuplicateCheck("user", "SSN", req.Ssn)
	if result == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "한명당 하나의 ID만 가질 수 있습니다",
		})
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ssn, err := bcrypt.GenerateFromPassword([]byte(req.Ssn), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	req.Pwd = string(pwd)
	req.Ssn = string(ssn)

	if req.Nickname != "" {
		_, err = db.NamedExec("INSERT INTO user(user_id, password, phone_number, SSN, name, nickname, agree) values (:user_id, :password, :phone_number, :SSN, :name, :nickname, :agree)", &req)
	} else {
		_, err = db.NamedExec("INSERT INTO user(user_id, password, phone_number, SSN, name, agree) values (:user_id, :password, :phone_number, :SSN, :name, :agree)", &req)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "계정 생성 성공!",
	})
}
