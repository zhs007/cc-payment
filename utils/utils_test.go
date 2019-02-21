package utils

import (
	"testing"

	"github.com/zhs007/cc-payment/proto"
)

func Test_ParseCurrencyString(t *testing.T) {
	type testData struct {
		currency string
		ret      paymentpb.Currency
	}

	lstTestData := []testData{
		testData{"USD", paymentpb.Currency_USD},
		testData{"EUR", paymentpb.Currency_EUR},
		testData{"CNY", paymentpb.Currency_NONECURRENCY},
	}

	for i := 0; i < len(lstTestData); i++ {
		ret := ParseCurrencyString(lstTestData[i].currency)
		if ret != lstTestData[i].ret {
			t.Fatalf("Test_ParseCurrencyString ParseCurrencyString(%v) return %v, need %v", lstTestData[i].currency, ret, lstTestData[i].ret)
		}
	}

	t.Logf("Test_ParseCurrencyString OK")
}
