package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/zhs007/cc-payment/proto"
)

// ParseCurrencyString - convert currency string to currency enumeration type
// 		if currency string is unavailable return NONECURRENCY
func ParseCurrencyString(currency string) paymentpb.Currency {
	c, isok := paymentpb.Currency_value[currency]
	if isok {
		return paymentpb.Currency(c)
	}

	return paymentpb.Currency_NONECURRENCY
}

// Currency2String - get currency name
func Currency2String(currency paymentpb.Currency) string {
	return paymentpb.Currency_name[int32(currency)]
}

// SendResponse - send a response with protobuf
func SendResponse(c *gin.Context, code int, pb proto.Message) {
	jsonstr, err := json.Marshal(pb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.String(code, string(jsonstr))
}
