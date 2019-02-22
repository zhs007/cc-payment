package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errdef "github.com/zhs007/cc-payment/err"
	"github.com/zhs007/cc-payment/model"
	"github.com/zhs007/cc-payment/proto"
	"github.com/zhs007/cc-payment/utils"
)

// Pay -
func Pay(c *gin.Context) {
	params := &paymentpb.PayParams{}
	c.BindJSON(params)

	currency := utils.ParseCurrencyString(params.Currency)
	if currency == paymentpb.Currency_NONECURRENCY {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errdef.ErrUnavailableCurrency,
		})

		return
	}

	if params.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errdef.ErrInvalidPaymentAmount,
		})

		return
	}

	payment, err := model.CreatePaymentWithUserID(params.Payer, params.Payee, params.Amount, currency)
	if err != nil {
		// StatusBadRequest
		if err == errdef.ErrUnavailableCurrency || err == errdef.ErrUnavailablePayer || err == errdef.ErrUnavailablePayee {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	payment, err = model.ApprovePayment(payment.PaymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	jsonstr, err := json.Marshal(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.String(http.StatusOK, string(jsonstr))
}

// GetPayments -
func GetPayments(c *gin.Context) {
	strstart := c.Query("start")
	strnums := c.Query("nums")
	strpayer := c.Query("payer")

	start := 0
	nums := 20
	payer := int64(0)

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

	if strpayer != "" {
		p, err := strconv.ParseInt(strpayer, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		payer = p
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errdef.ErrMissingParamsPayer.Error(),
		})
	}

	lst, err := model.GetPayments(payer, start, nums)
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
