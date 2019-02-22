package config

import (
	"testing"
)

func Test_loadTest1(t *testing.T) {
	err := load("../test/test1.yaml")
	if err != nil {
		t.Fatalf("Test_loadTest1 err %v", err)

		return
	}

	if cfg == nil {
		t.Fatalf("Test_loadTest1 cfg is nil")

		return
	}

	if cfg.PaymentDB.Host != "mysql" {
		t.Fatalf("Test_loadTest1 invalid PaymentDB.Host %v", cfg.PaymentDB.Host)
	}

	if cfg.PaymentDB.Port != 3306 {
		t.Fatalf("Test_loadTest1 invalid PaymentDB.Port %v", cfg.PaymentDB.Port)
	}

	if cfg.PaymentDB.User != "root" {
		t.Fatalf("Test_loadTest1 invalid PaymentDB.User %v", cfg.PaymentDB.User)
	}

	if cfg.PaymentDB.Password != "123456" {
		t.Fatalf("Test_loadTest1 invalid PaymentDB.Password %v", cfg.PaymentDB.Password)
	}

	if cfg.PaymentDB.Database != "ccpayment" {
		t.Fatalf("Test_loadTest1 invalid PaymentDB.Database %v", cfg.PaymentDB.Database)
	}

	if cfg.Log.LogLevel != "debug" {
		t.Fatalf("Test_loadTest1 invalid Log.LogLevel %v", cfg.Log.LogLevel)
	}

	if cfg.Log.OutputType != "console" {
		t.Fatalf("Test_loadTest1 invalid Log.OutputType %v", cfg.Log.OutputType)
	}

	if cfg.Log.LogPath != "./logs" {
		t.Fatalf("Test_loadTest1 invalid Log.LogPath %v", cfg.Log.LogPath)
	}

	if cfg.Service.Host != "127.0.0.1:8080" {
		t.Fatalf("Test_loadTest1 invalid Service.Host %v", cfg.Service.Host)
	}

	t.Logf("Test_loadTest1 OK")
}

func Test_loadTest2(t *testing.T) {
	err := load("../test/test2.yaml")
	if err != nil {
		t.Fatalf("Test_loadTest2 err %v", err)

		return
	}

	if cfg == nil {
		t.Fatalf("Test_loadTest2 cfg is nil")

		return
	}

	if cfg.PaymentDB.Host != "127.0.0.1" {
		t.Fatalf("Test_loadTest2 invalid PaymentDB.Host %v", cfg.PaymentDB.Host)
	}

	if cfg.PaymentDB.Port != 3307 {
		t.Fatalf("Test_loadTest2 invalid PaymentDB.Port %v", cfg.PaymentDB.Port)
	}

	if cfg.PaymentDB.User != "root" {
		t.Fatalf("Test_loadTest2 invalid PaymentDB.User %v", cfg.PaymentDB.User)
	}

	if cfg.PaymentDB.Password != "" {
		t.Fatalf("Test_loadTest2 invalid PaymentDB.Password %v", cfg.PaymentDB.Password)
	}

	if cfg.PaymentDB.Database != "ccpayment" {
		t.Fatalf("Test_loadTest2 invalid PaymentDB.Database %v", cfg.PaymentDB.Database)
	}

	if cfg.Log.LogLevel != "debug" {
		t.Fatalf("Test_loadTest2 invalid Log.LogLevel %v", cfg.Log.LogLevel)
	}

	if cfg.Log.OutputType != "console" {
		t.Fatalf("Test_loadTest2 invalid Log.OutputType %v", cfg.Log.OutputType)
	}

	if cfg.Log.LogPath != "./logs" {
		t.Fatalf("Test_loadTest2 invalid Log.LogPath %v", cfg.Log.LogPath)
	}

	if cfg.Service.Host != "127.0.0.1:8080" {
		t.Fatalf("Test_loadTest2 invalid Service.Host %v", cfg.Service.Host)
	}

	t.Logf("Test_loadTest2 OK")
}

func Test_isValidLogLevel(t *testing.T) {

	type testData struct {
		loglevel string
		ret      bool
	}

	lstTestData := []testData{
		testData{"debug", true},
		testData{"info", true},
		testData{"warn", true},
		testData{"error", true},
		testData{"", false},
		testData{"debug1", false},
		testData{"debug ", false},
		testData{" debug", false},
	}

	for i := 0; i < len(lstTestData); i++ {
		ret := isValidLogLevel(lstTestData[i].loglevel)
		if ret != lstTestData[i].ret {
			t.Fatalf("Test_isValidLogLevel isValidLogLevel(%v) return %v, need %v", lstTestData[i].loglevel, ret, lstTestData[i].ret)
		}
	}

	t.Logf("Test_isValidLogLevel OK")
}

func Test_isValidLogOutputType(t *testing.T) {

	type testData struct {
		outputtype string
		ret        bool
	}

	lstTestData := []testData{
		testData{"console", true},
		testData{"file", true},
		testData{"warn", false},
		testData{"error", false},
		testData{"", false},
		testData{"debug1", false},
		testData{"debug ", false},
		testData{" debug", false},
	}

	for i := 0; i < len(lstTestData); i++ {
		ret := isValidLogOutputType(lstTestData[i].outputtype)
		if ret != lstTestData[i].ret {
			t.Fatalf("Test_isValidLogOutputType isValidLogOutputType(%v) return %v, need %v", lstTestData[i].outputtype, ret, lstTestData[i].ret)
		}
	}

	t.Logf("Test_isValidLogOutputType OK")
}
