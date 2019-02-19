package utils

import (
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
