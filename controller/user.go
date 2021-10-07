package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	db, err := model.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't not connect to database",
		})
		return
	}
	defer db.Close()

	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Panic(err)
	}

	var jsonData map[string]string
	err = json.Unmarshal(value, &jsonData)
	if err != nil {
		log.Panic(err)
	}

	result := lib.DuplicateCheck("user", "user_id", jsonData["user_id"])

	if result != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID 중복됨",
		})
		return
	}

	var nickname sql.NullString
	if jsonData["nickname"] == "" {
		nickname.Valid = false
	} else {
		nickname.String = jsonData["nickname"]
		nickname.Valid = true
	}

	user := model.User{
		UserId:      jsonData["user_id"],
		Password:    model.Hash(jsonData["password"]),
		PhoneNumber: jsonData["phone_number"],
		SSN:         model.Hash(jsonData["SSN"]),
		Name:        jsonData["name"],
		Nickname:    nickname,
		Agree:       jsonData["agree"],
		UserType:    "general",
	}

	err = user.SignIn()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "계정 생성 성공!",
	})
}

func LoginHandler(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "바디 에러",
		})
	}
	var jsonData map[string]string
	err = json.Unmarshal(value, &jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "언마샬링 에러",
		})
		return
	}

	id := jsonData["user_id"]
	pwd := jsonData["password"]

	var result bool
	result, err = model.User{}.LoginWithID(id, pwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result == false {
		c.JSON(http.StatusOK, gin.H{
			"error": "not match id or password",
		})
	}

	c.JSON(200, true)
}

func OpenAccountHandler(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var value map[string]interface{}

	json.Unmarshal(body, &value)

	c.JSON(200, value)
}
