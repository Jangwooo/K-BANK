package controller

import (
	"crypto/sha512"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"K-BANK/data"
	"K-BANK/lib"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	db, err := data.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't not connect to database",
		})
		log.Panic()
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

	user := data.User{
		UserId:      jsonData["user_id"],
		Password:    fmt.Sprintf("%x", sha512.Sum512([]byte(jsonData["password"]))),
		PhoneNumber: jsonData["phone_number"],
		SSN:         jsonData["SSN"],
		Name:        jsonData["name"],
		Nickname:    nickname,
		Agree:       jsonData["agree"],
		UserType:    "general",
	}

	err = user.SignIn()

	if err != nil {
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massege": "계정 생성 성공!",
	})
}
