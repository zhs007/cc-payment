package api

import (
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
		utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
			Error: errdef.ErrUnavailableCurrency.Error(),
		})

		return
	}

	if params.Amount <= 0 {
		utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
			Error: errdef.ErrInvalidPaymentAmount.Error(),
		})

		return
	}

	payment, err := model.CreatePaymentWithUserID(params.Payer, params.Payee, params.Amount, currency)
	if err != nil {
		// StatusBadRequest
		if err == errdef.ErrUnavailableCurrency || err == errdef.ErrUnavailablePayer || err == errdef.ErrUnavailablePayee {
			utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
				Error: err.Error(),
			})

			return
		}

		utils.SendResponse(c, http.StatusInternalServerError, &paymentpb.ErrorResult{
			Error: err.Error(),
		})

		return
	}

	payment, err = model.ApprovePayment(payment.PaymentID)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, &paymentpb.ErrorResult{
			Error: err.Error(),
		})

		return
	}

	utils.SendResponse(c, http.StatusOK, payment)
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

		if n > 0 {
			nums = n
		}
	}

	if strpayer != "" {
		p, err := strconv.ParseInt(strpayer, 10, 64)
		if err != nil {
			utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
				Error: err.Error(),
			})

			return
		}

		payer = p
	} else {
		utils.SendResponse(c, http.StatusBadRequest, &paymentpb.ErrorResult{
			Error: errdef.ErrMissingParamsPayer.Error(),
		})

		return
	}

	lst, err := model.GetPayments(payer, start, nums)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, &paymentpb.ErrorResult{
			Error: err.Error(),
		})

		return
	}

	utils.SendResponse(c, http.StatusOK, lst)
}
