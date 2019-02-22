package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"

	"github.com/zhs007/cc-payment/config"
	"github.com/zhs007/cc-payment/logger"
	"github.com/zhs007/cc-payment/proto"
)

func onTestAccounts(t *testing.T, s *Serv) error {
	data := &paymentpb.UserList{
		TotalNums: 6,
		Users: []*paymentpb.User{
			&paymentpb.User{
				UserID:   10,
				UserName: "payeef",
				Status:   paymentpb.UserStatus_CANCOLLECT,
			},
			&paymentpb.User{
				UserID:   9,
				UserName: "payeee",
				Status:   paymentpb.UserStatus_CANCOLLECT,
			},
			&paymentpb.User{
				UserID:   7,
				UserName: "payeec",
				Status:   paymentpb.UserStatus_CANCOLLECT,
			},
			&paymentpb.User{
				UserID:   6,
				UserName: "payeeb",
				Status:   paymentpb.UserStatus_CANPAY,
			},
			&paymentpb.User{
				UserID:   3,
				UserName: "payerc",
				Status:   paymentpb.UserStatus_CANCOLLECT,
			},
			&paymentpb.User{
				UserID:   2,
				UserName: "payerb",
				Status:   paymentpb.UserStatus_CANPAY,
			},
		},
	}

	client := &http.Client{}
	resp, err := client.Get("http://127.0.0.1:8080/api/accounts")
	if err != nil {
		t.Fatalf("onTestAccounts Get %v", err)

		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("onTestAccounts ReadAll %v", err)

		return err
	}

	if resp.StatusCode != 200 {
		t.Fatalf("onTestAccounts Get StatusCode %v need 200", resp.StatusCode)

		return errors.New("onTestAccounts StatusCode err")
	}

	lst := &paymentpb.UserList{}
	err = jsonpb.UnmarshalString(string(body), lst)
	if err != nil {
		t.Fatalf("onTestAccounts UnmarshalString %v", err)

		return err
	}

	if lst.TotalNums != data.TotalNums {
		t.Fatalf("onTestAccounts totalNums err %v need %v", lst.TotalNums, data.TotalNums)

		return err
	}

	for i := 0; i < int(lst.TotalNums); i++ {
		if lst.Users[i].UserID != data.Users[i].UserID {
			t.Fatalf("onTestAccounts UserID err %v need %v", lst.Users[i].UserID, data.Users[i].UserID)

			return errors.New("onTestAccounts UserID err")
		}

		if lst.Users[i].UserName != data.Users[i].UserName {
			t.Fatalf("onTestAccounts UserName err %v need %v", lst.Users[i].UserName, data.Users[i].UserName)

			return errors.New("onTestAccounts UserName err")
		}

		if lst.Users[i].Status != data.Users[i].Status {
			t.Fatalf("onTestAccounts Status err %v need %v", lst.Users[i].Status, data.Users[i].Status)

			return errors.New("onTestAccounts Status err")
		}
	}

	fmt.Printf("%v\n", string(body))

	return nil
}

func onTestPay(t *testing.T, s *Serv) error {
	data := &paymentpb.UserPayment{
		Payer:    int64(3),
		Payee:    int64(7),
		Currency: paymentpb.Currency_USD,
		Amount:   int64(100),
		Status:   paymentpb.PaymentStatus_APPROVED,
	}

	params := paymentpb.PayParams{
		Payer:    data.Payer,
		Payee:    data.Payee,
		Currency: paymentpb.Currency_name[int32(data.Currency)],
		Amount:   data.Amount,
	}

	jsonparams, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("onTestPay Marshal %v", err)

		return err
	}

	client := &http.Client{}
	resp, err := client.Post("http://127.0.0.1:8080/api/pay", "application/json", bytes.NewReader(jsonparams))
	if err != nil {
		t.Fatalf("onTestPay Get %v", err)

		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("onTestPay ReadAll %v", err)

		return err
	}

	if resp.StatusCode != 200 {
		t.Fatalf("onTestPay Get StatusCode %v need 200", resp.StatusCode)

		return errors.New("onTestPay StatusCode err")
	}

	payment := &paymentpb.UserPayment{}
	err = jsonpb.UnmarshalString(string(body), payment)
	if err != nil {
		t.Fatalf("onTestPay UnmarshalString %v", err)

		return err
	}

	if payment.Payer != data.Payer {
		t.Fatalf("onTestPay Payer err %v need %v", payment.Payer, data.Payer)

		return errors.New("onTestPay Payer err")
	}

	if payment.Payee != data.Payee {
		t.Fatalf("onTestPay Payee err %v need %v", payment.Payee, data.Payee)

		return errors.New("onTestPay Payee err")
	}

	if payment.Currency != data.Currency {
		t.Fatalf("onTestPay Currency err %v need %v", payment.Currency, data.Currency)

		return errors.New("onTestPay Currency err")
	}

	if payment.Amount != data.Amount {
		t.Fatalf("onTestPay Amount err %v need %v", payment.Amount, data.Amount)

		return errors.New("onTestPay Amount err")
	}

	if payment.Status != data.Status {
		t.Fatalf("onTestPay Status err %v need %v", payment.Status, data.Status)

		return errors.New("onTestPay Status err")
	}

	fmt.Printf("%v\n", string(body))

	return nil
}

func onTestPayments(t *testing.T, s *Serv) error {
	// data := &paymentpb.UserList{
	// 	TotalNums: 6,
	// 	Users: []*paymentpb.User{
	// 		&paymentpb.User{
	// 			UserID:   10,
	// 			UserName: "payeef",
	// 			Status:   paymentpb.UserStatus_CANCOLLECT,
	// 		},
	// 		&paymentpb.User{
	// 			UserID:   9,
	// 			UserName: "payeee",
	// 			Status:   paymentpb.UserStatus_CANCOLLECT,
	// 		},
	// 		&paymentpb.User{
	// 			UserID:   7,
	// 			UserName: "payeec",
	// 			Status:   paymentpb.UserStatus_CANCOLLECT,
	// 		},
	// 		&paymentpb.User{
	// 			UserID:   6,
	// 			UserName: "payeeb",
	// 			Status:   paymentpb.UserStatus_CANPAY,
	// 		},
	// 		&paymentpb.User{
	// 			UserID:   3,
	// 			UserName: "payerc",
	// 			Status:   paymentpb.UserStatus_CANCOLLECT,
	// 		},
	// 		&paymentpb.User{
	// 			UserID:   2,
	// 			UserName: "payerb",
	// 			Status:   paymentpb.UserStatus_CANPAY,
	// 		},
	// 	},
	// }

	client := &http.Client{}
	resp, err := client.Get("http://127.0.0.1:8080/api/payments?payer=3")
	if err != nil {
		t.Fatalf("onTestPayments Get %v", err)

		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("onTestPayments ReadAll %v", err)

		return err
	}

	fmt.Printf("%v\n", string(body))

	if resp.StatusCode != 200 {
		t.Fatalf("onTestPayments Get StatusCode %v need 200", resp.StatusCode)

		return errors.New("onTestPayments StatusCode err")
	}

	// lst := &paymentpb.UserList{}
	// err = jsonpb.UnmarshalString(string(body), lst)
	// if err != nil {
	// 	t.Fatalf("onTestAccounts UnmarshalString %v", err)

	// 	return err
	// }

	// if lst.TotalNums != data.TotalNums {
	// 	t.Fatalf("onTestAccounts totalNums err %v need %v", lst.TotalNums, data.TotalNums)

	// 	return err
	// }

	// for i := 0; i < int(lst.TotalNums); i++ {
	// 	if lst.Users[i].UserID != data.Users[i].UserID {
	// 		t.Fatalf("onTestAccounts UserID err %v need %v", lst.Users[i].UserID, data.Users[i].UserID)

	// 		return errors.New("onTestAccounts UserID err")
	// 	}

	// 	if lst.Users[i].UserName != data.Users[i].UserName {
	// 		t.Fatalf("onTestAccounts UserName err %v need %v", lst.Users[i].UserName, data.Users[i].UserName)

	// 		return errors.New("onTestAccounts UserName err")
	// 	}

	// 	if lst.Users[i].Status != data.Users[i].Status {
	// 		t.Fatalf("onTestAccounts Status err %v need %v", lst.Users[i].Status, data.Users[i].Status)

	// 		return errors.New("onTestAccounts Status err")
	// 	}
	// }

	return nil
}

func Test_Serv(t *testing.T) {
	err := config.LoadConfig("./test/testdocker.yaml")
	if err != nil {
		t.Fatalf("Test_Serv LoadConfig %v", err)

		return
	}

	err = logger.InitLogger()
	if err != nil {
		t.Fatalf("Test_Serv InitLogger %v", err)

		return
	}

	cfg, isok := config.GetConfig()
	if !isok {
		t.Fatalf("Test_Serv GetConfig %v", err)

		return
	}

	s := StartServ(cfg.Service.Host)

	//!!! wait serv started
	time.Sleep(1 * time.Second)

	go func() {
		err = onTestAccounts(t, s)
		if err != nil {
			s.Stop()

			return
		}

		err = onTestPay(t, s)
		if err != nil {
			s.Stop()

			return
		}

		err = onTestPayments(t, s)
		if err != nil {
			s.Stop()

			return
		}

		s.Stop()
	}()

	// time.Sleep(10 * time.Second)
	s.Wait()

	if err != nil {
		t.Fatalf("Test_Serv err %v", err)

		return
	}

	t.Logf("Test_Serv OK")
}
