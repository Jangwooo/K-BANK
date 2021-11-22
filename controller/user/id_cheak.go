package user

import (
	"net/http"

	"K-BANK/lib"
	"github.com/gin-gonic/gin"
)

func IdCheck(c *gin.Context) {
	k := c.Query("key")
	v := c.Query("value")
	result := lib.DuplicateCheck(k, v)

	if result {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}

}
