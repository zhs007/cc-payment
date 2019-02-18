package model

import (
	"database/sql"
	"fmt"

	// import mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/zhs007/cc-payment/config"
	"github.com/zhs007/cc-payment/err"
)

type paymentDB struct {
	db *sql.DB
}

func newPaymentDB() (*paymentDB, error) {
	cfg, isok := config.GetConfig()
	if !isok {
		return nil, err.ErrNotLoadConfig
	}
	dsn := fmt.Sprintf("%v:%v", cfg.PaymentDB.User, cfg.PaymentDB.Password)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &paymentDB{
		db: db,
	}, nil
}
