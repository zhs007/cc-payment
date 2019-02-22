package model

import (
	"github.com/zhs007/cc-payment/proto"
)

// CreatePaymentWithUserID - create payment with userid
func CreatePaymentWithUserID(payer int64, payee int64, amount int64,
	currency paymentpb.Currency) (*paymentpb.UserPayment, error) {

	pdb, err := getPaymentDB()
	if err != nil {
		return nil, err
	}

	return pdb.createPayment(payer, payee, amount, currency)
}

// ApprovePayment - approve payment
func ApprovePayment(paymentid int64) (*paymentpb.UserPayment, error) {

	pdb, err := getPaymentDB()
	if err != nil {
		return nil, err
	}

	return pdb.approvePayment(paymentid)
}
