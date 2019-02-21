package model

import (
	"testing"

	"github.com/zhs007/cc-payment/proto"

	"github.com/zhs007/cc-payment/config"
)

func Test_getAccountList(t *testing.T) {
	config.LoadConfig("../test/testdocker.yaml")

	pdb, err := newPaymentDB()
	if err != nil {
		t.Fatalf("Test_getAccountList newPaymentDB() err %v", err)

		return
	}

	// make result data
	mapUsers := make(map[string]paymentpb.UserStatus)
	mapUsers["payera"] = paymentpb.UserStatus_NORMALUSER
	mapUsers["payerb"] = paymentpb.UserStatus_CANPAY
	mapUsers["payerc"] = paymentpb.UserStatus_CANCOLLECT
	mapUsers["payerd"] = paymentpb.UserStatus_FROZEN
	mapUsers["payeea"] = paymentpb.UserStatus_NORMALUSER
	mapUsers["payeeb"] = paymentpb.UserStatus_CANPAY
	mapUsers["payeec"] = paymentpb.UserStatus_CANCOLLECT
	mapUsers["payeed"] = paymentpb.UserStatus_FROZEN

	lst, err := pdb.getAccountList(0, 100)
	if err != nil {
		t.Fatalf("Test_getAccountList getAccountList() err %v", err)

		return
	}

	// only 4 valid user
	if len(lst.Users) != 4 {
		t.Fatalf("Test_getAccountList getAccountList() invalid user nums %v", len(lst.Users))

		return
	}

	for i := 0; i < len(lst.Users); i++ {
		curuser := lst.Users[i]

		// check username
		curstatus, isok := mapUsers[curuser.UserName]
		if !isok {
			t.Fatalf("Test_getAccountList getAccountList() non user %v", curuser)
		}

		// check status
		if curstatus != curuser.Status {
			t.Fatalf("Test_getAccountList getAccountList() user status err get %v - need %v", curuser.Status, curstatus)
		}
	}

	t.Logf("Test_getAccountList OK")
}
