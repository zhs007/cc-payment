package model

import (
	"testing"

	"github.com/zhs007/cc-payment/config"
	errdef "github.com/zhs007/cc-payment/err"
	"github.com/zhs007/cc-payment/proto"
)

func Test_getAccountList(t *testing.T) {
	config.LoadConfig("../test/testdocker.yaml")

	pdb, err := newPaymentDB()
	if err != nil {
		t.Fatalf("Test_getAccountList newPaymentDB() err %v", err)

		return
	}

	// make result data
	mapUsers := make(map[string]paymentpb.User)
	mapUsers["payera"] = paymentpb.User{Status: paymentpb.UserStatus_NORMALUSER, UserID: 1}
	mapUsers["payerb"] = paymentpb.User{Status: paymentpb.UserStatus_CANPAY, UserID: 2}
	mapUsers["payerc"] = paymentpb.User{Status: paymentpb.UserStatus_CANCOLLECT, UserID: 3}
	mapUsers["payerd"] = paymentpb.User{Status: paymentpb.UserStatus_FROZEN, UserID: 4}
	mapUsers["payeea"] = paymentpb.User{Status: paymentpb.UserStatus_NORMALUSER, UserID: 5}
	mapUsers["payeeb"] = paymentpb.User{Status: paymentpb.UserStatus_CANPAY, UserID: 6}
	mapUsers["payeec"] = paymentpb.User{Status: paymentpb.UserStatus_CANCOLLECT, UserID: 7}
	mapUsers["payeed"] = paymentpb.User{Status: paymentpb.UserStatus_FROZEN, UserID: 8}
	mapUsers["payeee"] = paymentpb.User{Status: paymentpb.UserStatus_CANCOLLECT, UserID: 9}
	mapUsers["payeef"] = paymentpb.User{Status: paymentpb.UserStatus_CANCOLLECT, UserID: 10}

	lst, err := pdb.getAccountList(0, 100)
	if err != nil {
		t.Fatalf("Test_getAccountList getAccountList() err %v", err)

		return
	}

	// only 6 valid user
	if len(lst.Users) != 6 {
		t.Fatalf("Test_getAccountList getAccountList() invalid user nums %v", len(lst.Users))

		return
	}

	for i := 0; i < len(lst.Users); i++ {
		curuser := lst.Users[i]

		// check username
		curuser0, isok := mapUsers[curuser.UserName]
		if !isok {
			t.Fatalf("Test_getAccountList getAccountList() non user %v", curuser)
		}

		// check status
		if curuser0.Status != curuser.Status {
			t.Fatalf("Test_getAccountList getAccountList() user status err get %v - need %v", curuser.Status, curuser0.Status)
		}

		// check userid
		if curuser0.UserID != curuser.UserID {
			t.Fatalf("Test_getAccountList getAccountList() userid err get %v - need %v", curuser.UserID, curuser0.UserID)
		}
	}

	t.Logf("Test_getAccountList OK")
}

func Test_Payment(t *testing.T) {
	config.LoadConfig("../test/testdocker.yaml")

	pdb, err := newPaymentDB()
	if err != nil {
		t.Fatalf("Test_Payment newPaymentDB() err %v", err)

		return
	}

	// 2 -> 7 => ok
	payment, err := pdb.createPayment(2, 7, 5000, paymentpb.Currency_USD)
	if err != nil {
		t.Fatalf("Test_Payment createPayment(ok) err %v", err)

		return
	}

	if payment.StartBalancePayer != 10000 {
		t.Fatalf("Test_Payment createPayment(ok) StartBalancePayer err %v", payment.StartBalancePayer)

		return
	}

	// ErrExistPayment
	_, err = pdb.createPayment(2, 7, 5000, paymentpb.Currency_USD)
	if err == nil || err != errdef.ErrExistPayment {
		t.Fatalf("Test_Payment createPayment(ErrExistPayment) err %v", err)

		return
	}

	payment, err = pdb.approvePayment(payment.PaymentID)
	if err != nil {
		t.Fatalf("Test_Payment approvePayment(ok) err %v", err)

		return
	}

	// 2 -> 7 => no EUR
	_, err = pdb.createPayment(2, 7, 5000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrUnavailablePayerCurrency {
		t.Fatalf("Test_Payment createPayment(no EUR) err %v", err)

		return
	}

	// 2 -> 7 => ok2
	payment, err = pdb.createPayment(2, 7, 5000, paymentpb.Currency_USD)
	if err != nil {
		t.Fatalf("Test_Payment createPayment(ok2) err %v", err)

		return
	}

	if payment.StartBalancePayer != 5000 {
		t.Fatalf("Test_Payment createPayment(ok2) StartBalancePayer err %v", payment.StartBalancePayer)

		return
	}

	payment, err = pdb.approvePayment(payment.PaymentID)
	if err != nil {
		t.Fatalf("Test_Payment approvePayment(ok2) err %v", err)

		return
	}

	// 2 -> 7 => ErrInsufficientBalance
	payment, err = pdb.createPayment(2, 7, 5000, paymentpb.Currency_USD)
	if err == nil || err != errdef.ErrInsufficientBalance {
		t.Fatalf("Test_Payment createPayment(ErrInsufficientBalance) err %v", err)

		return
	}

	// 3 -> 10 => ErrUnavailablePayeeCurrency (no currency)
	payment, err = pdb.createPayment(3, 10, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrUnavailablePayeeCurrency {
		t.Fatalf("Test_Payment createPayment(ErrUnavailablePayeeCurrency) err %v", err)

		return
	}

	// 3 -> 9 => ErrUnavailablePayeeCurrency (invalid currency)
	payment, err = pdb.createPayment(3, 10, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrUnavailablePayeeCurrency {
		t.Fatalf("Test_Payment createPayment(ErrUnavailablePayeeCurrency) err %v", err)

		return
	}

	// 1 -> 7 => ErrCannotPay
	payment, err = pdb.createPayment(1, 7, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrCannotPay {
		t.Fatalf("Test_Payment createPayment(ErrCannotPay) err %v", err)

		return
	}

	// 4 -> 7 => ErrFrozenPayer
	payment, err = pdb.createPayment(4, 7, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrFrozenPayer {
		t.Fatalf("Test_Payment createPayment(ErrFrozenPayer) err %v", err)

		return
	}

	// 3 -> 5 => ErrCannotCollect
	payment, err = pdb.createPayment(3, 5, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrCannotCollect {
		t.Fatalf("Test_Payment createPayment(ErrCannotCollect) err %v", err)

		return
	}

	// 3 -> 8 => ErrFrozenPayee
	payment, err = pdb.createPayment(3, 8, 1000, paymentpb.Currency_EUR)
	if err == nil || err != errdef.ErrFrozenPayee {
		t.Fatalf("Test_Payment createPayment(ErrFrozenPayee) err %v", err)

		return
	}

	// ErrPaymentApproved
	payment, err = pdb.approvePayment(1)
	if err == nil || err != errdef.ErrPaymentApproved {
		t.Fatalf("Test_Payment approvePayment(ErrPaymentApproved) err %v", err)

		return
	}

	// ErrNoPayment
	payment, err = pdb.approvePayment(10)
	if err == nil || err != errdef.ErrNoPayment {
		t.Fatalf("Test_Payment approvePayment(ErrNoPayment) err %v", err)

		return
	}

	t.Logf("Test_Payment OK")
}
