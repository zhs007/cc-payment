package api

import (
	"net/http"
	"strconv"

	"github.com/zhs007/cc-payment/proto"
	"github.com/zhs007/cc-payment/utils"

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
			utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
				Error: err.Error(),
			})

			return
		}

		if s >= 0 {
			start = s
		}
	}

	if strnums != "" {
		n, err := strconv.Atoi(strnums)
		if err != nil {
			utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
				Error: err.Error(),
			})

			return
		}

		if nums > 0 {
			nums = n
		}
	}

	lst, err := model.GetAccounts(start, nums)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, &paymentpb.ErrorResult{
			Error: err.Error(),
		})

		return
	}

	utils.SendResponse(c, http.StatusOK, lst)
}
