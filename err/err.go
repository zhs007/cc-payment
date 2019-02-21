package err

import "errors"

var (
	// ErrConfigLogOutputType - invalid config.yaml:log.outputtype
	ErrConfigLogOutputType = errors.New("invalid config.yaml:log.outputtype")
	// ErrConfigLogLevel - invalid config.yaml:log.loglevel
	ErrConfigLogLevel = errors.New("invalid config.yaml:log.loglevel")
	// ErrNotLoadConfig - not load config.yaml
	ErrNotLoadConfig = errors.New("not load config.yaml")

	// ErrUnavailablePaymentDB - the payment database is not available
	ErrUnavailablePaymentDB = errors.New("the payment database is not available")

	// ErrUnavailableCurrency - unavailable currency
	ErrUnavailableCurrency = errors.New("unavailable currency")
	// ErrUnavailablePayer - unavailable payer
	ErrUnavailablePayer = errors.New("unavailable payer")
	// ErrUnavailablePayee - unavailable payee
	ErrUnavailablePayee = errors.New("unavailable payee")
	// ErrFrozenPayer - payer is a frozen account
	ErrFrozenPayer = errors.New("payer is a frozen account")
	// ErrFrozenPayee - payee is a frozen account
	ErrFrozenPayee = errors.New("payee is a frozen account")
	// ErrUnavailablePayerCurrency - payer has not the currency
	ErrUnavailablePayerCurrency = errors.New("payer has not the currency")
	// ErrUnavailablePayeeCurrency - payee has not the currency
	ErrUnavailablePayeeCurrency = errors.New("payee has not the currency")
	// ErrInsufficientBalance - insufficient balance
	ErrInsufficientBalance = errors.New("insufficient balance")
	// ErrCannotPay - this account cannot pay
	ErrCannotPay = errors.New("this account cannot pay")
	// ErrCannotCollect - this account cannot collect
	ErrCannotCollect = errors.New("this account cannot collect")
	// ErrPaymentApproved - the payment approved
	ErrPaymentApproved = errors.New("the payment approved")
	// ErrPaymentFailed - the payment failed
	ErrPaymentFailed = errors.New("the payment failed")
	// ErrExistPayment - other payments that already exist
	ErrExistPayment = errors.New("other payments that already exist")
	// ErrNoPayment - no payment
	ErrNoPayment = errors.New("no payment")

	// ErrInvalidPaymentAmount - invalid payment amount
	ErrInvalidPaymentAmount = errors.New("invalid payment amount")

	// ErrInvalidRowsAffected - invalid RowsAffected
	ErrInvalidRowsAffected = errors.New("invalid RowsAffected")
)
