package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/zhs007/cc-payment/model"
)

// GetAccounts -
func GetAccounts(c *gin.Context) {
	strstart := c.Query("start")
	strnums := c.Query("nums")

	start := 0
	nums := 20

	if strstart != "" {
		s, err := strconv.Atoi(strstart)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		start = s
	}

	if strnums != "" {
		n, err := strconv.Atoi(strnums)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		nums = n
	}

	lst, err := model.GetAccounts(start, nums)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	jsonstr, err := json.Marshal(lst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.String(http.StatusOK, string(jsonstr))
}
