package model

import (
	"github.com/zhs007/cc-payment/proto"
)

// GetAccounts - get accounts
func GetAccounts(start int, nums int) (*paymentpb.UserList, error) {
	pdb, err := getPaymentDB()
	if err != nil {
		return nil, err
	}

	return pdb.getAccountList(start, nums)
}
