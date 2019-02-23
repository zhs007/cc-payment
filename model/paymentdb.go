package model

import (
	"database/sql"
	"fmt"
	"sync"

	// import mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/zhs007/cc-payment/config"
	errdef "github.com/zhs007/cc-payment/err"
	"github.com/zhs007/cc-payment/proto"
	"github.com/zhs007/cc-payment/utils"
)

// paymentDB - payment database
type paymentDB struct {
	db *sql.DB
}

var dbPayment *paymentDB
var onceDBPayment sync.Once

// getPaymentDB - get singleton paymentDB
func getPaymentDB() (*paymentDB, error) {
	var err error

	onceDBPayment.Do(func() {
		dbPayment, err = newPaymentDB()
	})

	return dbPayment, err
}

// newPaymentDB - new paymentDB
func newPaymentDB() (*paymentDB, error) {
	cfg, isok := config.GetConfig()
	if !isok {
		return nil, errdef.ErrNotLoadConfig
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.PaymentDB.User, cfg.PaymentDB.Password, cfg.PaymentDB.Host, cfg.PaymentDB.Port, cfg.PaymentDB.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &paymentDB{
		db: db,
	}, nil
}

// _countAccountList - count the number of available accounts
func (pdb *paymentDB) _countAccountList() (int, error) {
	totalnums := 0
	row := pdb.db.QueryRow("select count(userid) from users where status > 0 and status < 3 ")
	err := row.Scan(&totalnums)
	if err != nil {
		return -1, err
	}

	return totalnums, nil
}

// getAccountList - get available accounts
func (pdb *paymentDB) getAccountList(start int, nums int) (*paymentpb.UserList, error) {
	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	totalnums, err := pdb._countAccountList()
	if err != nil {
		return nil, err
	}

	// if the index overflows
	if start >= totalnums {
		return &paymentpb.UserList{
			TotalNums:  int32(totalnums),
			StartIndex: int32(start),
			PageNums:   0,
		}, nil
	}

	rows, err := pdb.db.Query("select userid, username, status, UNIX_TIMESTAMP(registertime) as registertime from users where status > 0 and status < 3 order by userid desc limit ?, ?",
		start, nums)

	if err != nil {
		if err == sql.ErrNoRows {
			return &paymentpb.UserList{}, nil
		}

		return nil, err
	}
	defer rows.Close()

	lstuser := &paymentpb.UserList{
		StartIndex: int32(start),
		PageNums:   0,
		TotalNums:  int32(totalnums),
	}

	for rows.Next() {
		curuser := &paymentpb.User{}
		err := rows.Scan(&curuser.UserID, &curuser.UserName, &curuser.Status, &curuser.RegisterTime)
		if err != nil {
			return lstuser, err
		}

		lstuser.Users = append(lstuser.Users, curuser)
		lstuser.PageNums++
	}

	err = rows.Err()
	if err != nil {
		return lstuser, err
	}

	return lstuser, nil
}

// getUserCurrencies - get user currencies
func (pdb *paymentDB) getUserCurrencies(userid int64) (*paymentpb.UserCurrencies, error) {
	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	rows, err := pdb.db.Query("select currency, balance from usercurrencies where userid = ?",
		userid)

	if err != nil {
		if err == sql.ErrNoRows {
			return &paymentpb.UserCurrencies{}, nil
		}

		return nil, err
	}
	defer rows.Close()

	lst := &paymentpb.UserCurrencies{
		Currencies: make(map[string]*paymentpb.UserCurrency),
	}

	for rows.Next() {
		currency := &paymentpb.UserCurrency{}
		err := rows.Scan(&currency.CurrencyString, &currency.Balance)
		if err != nil {
			return lst, err
		}

		currency.Currency = utils.ParseCurrencyString(currency.CurrencyString)
		if currency.Currency == paymentpb.Currency_NONECURRENCY {
			return lst, err
		}

		lst.Currencies[currency.CurrencyString] = currency
	}

	err = rows.Err()
	if err != nil {
		return lst, err
	}

	return lst, nil
}

// _checkPayerOnCreate - check payer
func (pdb *paymentDB) _checkPayerOnCreate(payer int64, amount int64, currency paymentpb.Currency) (int64, error) {
	if amount <= 0 {
		return int64(-1), errdef.ErrInvalidPaymentAmount
	}

	status := 0
	err := pdb.db.QueryRow("select status from users where userid = ?", payer).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return int64(-1), errdef.ErrUnavailablePayer
		}

		return int64(-1), err
	}

	if status == int(paymentpb.UserStatus_FROZEN) {
		return int64(-1), errdef.ErrFrozenPayer
	}

	if status != int(paymentpb.UserStatus_CANPAY) &&
		status != int(paymentpb.UserStatus_CANCOLLECT) {

		return int64(-1), errdef.ErrCannotPay
	}

	var balance int64
	var frozen int64
	err = pdb.db.QueryRow("select balance, frozen from usercurrencies where userid = ? and currency = ?",
		payer, utils.Currency2String(currency)).Scan(&balance, &frozen)
	if err != nil {
		if err == sql.ErrNoRows {
			return int64(-1), errdef.ErrUnavailablePayerCurrency
		}

		return int64(-1), err
	}

	if frozen > 0 {
		return balance, errdef.ErrExistPayment
	}

	if balance < amount {
		return balance, errdef.ErrInsufficientBalance
	}

	return balance, nil
}

// _checkPayeeOnCreate - check payee
func (pdb *paymentDB) _checkPayeeOnCreate(payee int64, currency paymentpb.Currency) error {
	status := 0
	err := pdb.db.QueryRow("select status from users where userid = ?", payee).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return errdef.ErrUnavailablePayee
		}

		return err
	}

	if status == int(paymentpb.UserStatus_FROZEN) {
		return errdef.ErrFrozenPayee
	}

	if status != int(paymentpb.UserStatus_CANCOLLECT) {
		return errdef.ErrCannotCollect
	}

	var balance int64
	err = pdb.db.QueryRow("select balance from usercurrencies where userid = ? and currency = ? and balance >= 0",
		payee, utils.Currency2String(currency)).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return errdef.ErrUnavailablePayeeCurrency
		}

		return err
	}

	return nil
}

// createPayment - create payment
func (pdb *paymentDB) createPayment(payer int64, payee int64, amount int64,
	currency paymentpb.Currency) (*paymentpb.UserPayment, error) {

	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	oldbalance, err := pdb._checkPayerOnCreate(payer, amount, currency)
	if err != nil {
		return nil, err
	}

	err = pdb._checkPayeeOnCreate(payee, currency)
	if err != nil {
		return nil, err
	}

	tx, err := pdb.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := tx.Exec("update usercurrencies set frozen = ?, balance = balance - ? where userid = ? and currency = ? and frozen = 0 and balance >= ?",
		amount, amount, payer, utils.Currency2String(currency), amount)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		return nil, errdef.ErrInvalidRowsAffected
	}

	res, err = tx.Exec("insert into userpayments(payer, payee, currency, amount, startbalance0) values(?, ?, ?, ?, ?)",
		payer, payee, utils.Currency2String(currency), amount, oldbalance)
	if err != nil {
		return nil, err
	}

	paymentid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &paymentpb.UserPayment{
		PaymentID:         paymentid,
		Payer:             payer,
		Payee:             payee,
		Currency:          currency,
		Status:            paymentpb.PaymentStatus_CREATED,
		StartBalancePayer: oldbalance,
	}, nil
}

// _checkPayerOnApprove - check payer
func (pdb *paymentDB) _checkPayerOnApprove(payer int64, amount int64, currency paymentpb.Currency) error {
	if amount <= 0 {
		return errdef.ErrInvalidPaymentAmount
	}

	status := 0
	err := pdb.db.QueryRow("select status from users where userid = ?", payer).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return errdef.ErrUnavailablePayer
		}

		return err
	}

	if status == int(paymentpb.UserStatus_FROZEN) {
		return errdef.ErrFrozenPayer
	}

	if status != int(paymentpb.UserStatus_CANPAY) &&
		status != int(paymentpb.UserStatus_CANCOLLECT) {

		return errdef.ErrCannotPay
	}

	var frozen int64
	err = pdb.db.QueryRow("select frozen from usercurrencies where userid = ? and currency = ?",
		payer, utils.Currency2String(currency)).Scan(&frozen)
	if err != nil {
		if err == sql.ErrNoRows {
			return errdef.ErrUnavailablePayerCurrency
		}

		return err
	}

	if frozen != amount {
		return errdef.ErrInsufficientBalance
	}

	return nil
}

// _checkPayeeOnApprove - check payee
func (pdb *paymentDB) _checkPayeeOnApprove(payee int64, currency paymentpb.Currency) (int64, error) {
	status := 0
	err := pdb.db.QueryRow("select status from users where userid = ?", payee).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return int64(-1), errdef.ErrUnavailablePayer
		}

		return int64(-1), err
	}

	if status == int(paymentpb.UserStatus_FROZEN) {
		return int64(-1), errdef.ErrFrozenPayer
	}

	if status != int(paymentpb.UserStatus_CANCOLLECT) {
		return int64(-1), errdef.ErrCannotCollect
	}

	var balance int64
	err = pdb.db.QueryRow("select balance from usercurrencies where userid = ? and currency = ?",
		payee, utils.Currency2String(currency)).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return int64(-1), errdef.ErrUnavailablePayerCurrency
		}

		return int64(-1), err
	}

	return balance, nil
}

// _getUserBalance - get user currency balance
func (pdb *paymentDB) _getUserBalance(tx *sql.Tx, payee int64, currency paymentpb.Currency) (int64, error) {
	var balance int64
	err := tx.QueryRow("select balance from usercurrencies where userid = ? and currency = ?",
		payee, utils.Currency2String(currency)).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return int64(-1), errdef.ErrUnavailablePayerCurrency
		}

		return int64(-1), err
	}

	return balance, nil
}

// approvePayment - approve payment
func (pdb *paymentDB) approvePayment(paymentid int64) (*paymentpb.UserPayment, error) {

	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	var payer int64
	var payee int64
	var strcurrency string
	var amount int64
	var status int32
	err := pdb.db.QueryRow("select payer, payee, currency, amount, paymentstatus from userpayments where id = ?",
		paymentid).Scan(&payer, &payee, &strcurrency, &amount, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errdef.ErrNoPayment
		}

		return nil, err
	}

	if status == int32(paymentpb.PaymentStatus_APPROVED) {
		return nil, errdef.ErrPaymentApproved
	}

	if status == int32(paymentpb.PaymentStatus_FAILED) {
		return nil, errdef.ErrPaymentFailed
	}

	currency := utils.ParseCurrencyString(strcurrency)
	if currency == paymentpb.Currency_NONECURRENCY {
		return nil, errdef.ErrUnavailableCurrency
	}

	err = pdb._checkPayerOnApprove(payer, amount, currency)
	if err != nil {
		return nil, err
	}

	oldbalance, err := pdb._checkPayeeOnApprove(payee, currency)
	if err != nil {
		return nil, err
	}

	tx, err := pdb.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := tx.Exec("update usercurrencies set balance = balance + ? where userid = ? and currency = ?",
		amount, payee, utils.Currency2String(currency))
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		return nil, errdef.ErrInvalidRowsAffected
	}

	res, err = tx.Exec("update usercurrencies set frozen = 0 where userid = ? and currency = ? and frozen = ?",
		payer, utils.Currency2String(currency), amount)
	if err != nil {
		return nil, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		return nil, errdef.ErrInvalidRowsAffected
	}

	payerbalance, err := pdb._getUserBalance(tx, payer, currency)
	if err != nil {
		return nil, err
	}

	payeebalance, err := pdb._getUserBalance(tx, payee, currency)
	if err != nil {
		return nil, err
	}

	res, err = tx.Exec("update userpayments set startbalance1 = ?, endbalance0 = ?, endbalance1 = ?, paymentstatus = ?, donetime = unix_timestamp() where id = ?",
		oldbalance, payerbalance, payeebalance, paymentpb.PaymentStatus_APPROVED, paymentid)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &paymentpb.UserPayment{
		PaymentID: paymentid,
		Payer:     payer,
		Payee:     payee,
		Currency:  currency,
		Amount:    amount,
		Status:    paymentpb.PaymentStatus_APPROVED,
	}, nil
}

// cancelPayment - cancel payment
func (pdb *paymentDB) cancelPayment(paymentid int64) (*paymentpb.UserPayment, error) {

	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	var payer int64
	var strcurrency string
	var amount int64
	var status int32
	err := pdb.db.QueryRow("select payer, currency, amount, paymentstatus from userpayments where id = ?",
		paymentid).Scan(&payer, &strcurrency, &amount, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errdef.ErrNoPayment
		}

		return nil, err
	}

	if status == int32(paymentpb.PaymentStatus_APPROVED) {
		return nil, errdef.ErrPaymentApproved
	}

	if status == int32(paymentpb.PaymentStatus_FAILED) {
		return nil, errdef.ErrPaymentFailed
	}

	currency := utils.ParseCurrencyString(strcurrency)
	if currency == paymentpb.Currency_NONECURRENCY {
		return nil, errdef.ErrUnavailableCurrency
	}

	err = pdb._checkPayerOnApprove(payer, amount, currency)
	if err != nil {
		return nil, err
	}

	tx, err := pdb.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := tx.Exec("update usercurrencies set balance = balance + ?, frozen = 0 where userid = ? and currency = ? and frozen = ?",
		amount, payer, utils.Currency2String(currency), amount)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		return nil, errdef.ErrInvalidRowsAffected
	}

	res, err = tx.Exec("update userpayments set paymentstatus = ?, donetime = unix_timestamp() where id = ?",
		paymentpb.PaymentStatus_FAILED, paymentid)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// _countPayments - count the number of available payments
func (pdb *paymentDB) _countPayments(payer int64) (int, error) {
	totalnums := 0
	row := pdb.db.QueryRow("select count(id) from userpayments where payer = ? and paymentstatus = 2 order by id desc", payer)
	err := row.Scan(&totalnums)
	if err != nil {
		return -1, err
	}

	return totalnums, nil
}

// getPaymentList - get payments
func (pdb *paymentDB) getPaymentList(payer int64, start int, nums int) (*paymentpb.UserPayments, error) {

	if pdb.db == nil {
		return nil, errdef.ErrUnavailablePaymentDB
	}

	totalnums, err := pdb._countPayments(payer)
	if err != nil {
		return nil, err
	}

	if totalnums == 0 {
		return &paymentpb.UserPayments{
			TotalNums: int32(totalnums),
		}, nil
	}

	if start >= totalnums {
		return &paymentpb.UserPayments{
			TotalNums:  int32(totalnums),
			StartIndex: int32(totalnums),
			PageNums:   0,
		}, nil
	}

	rows, err := pdb.db.Query("select payer, payee, amount, unix_timestamp(donetime), currency from userpayments where payer = ? and paymentstatus = 2 order by id desc limit ?, ?",
		payer, start, nums)

	if err != nil {
		if err == sql.ErrNoRows {
			return &paymentpb.UserPayments{
				TotalNums: int32(totalnums),
			}, nil
		}

		return nil, err
	}
	defer rows.Close()

	lst := &paymentpb.UserPayments{
		TotalNums:  int32(totalnums),
		StartIndex: int32(start),
		PageNums:   0,
	}

	for rows.Next() {
		up := &paymentpb.UserPayment{}
		var curcurrency string
		err := rows.Scan(&up.Payer, &up.Payee, &up.Amount, &up.DoneTime, &curcurrency)
		if err != nil {
			return lst, err
		}

		up.Currency = utils.ParseCurrencyString(curcurrency)
		if up.Currency == paymentpb.Currency_NONECURRENCY {
			return lst, err
		}

		lst.Payments = append(lst.Payments, up)
	}

	err = rows.Err()
	if err != nil {
		return lst, err
	}

	return lst, nil
}
