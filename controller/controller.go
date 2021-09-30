package controller

import (
	"encoding/json"
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

}
