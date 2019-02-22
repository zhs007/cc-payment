package main

import (
	"fmt"
	"os"

	"github.com/zhs007/cc-payment/config"
	"github.com/zhs007/cc-payment/logger"
)

func main() {
	err := config.LoadConfig("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig fail! %v", err)

		os.Exit(-1)
	}

	err = logger.InitLogger()
	if err != nil {
		fmt.Printf("InitLogger fail! %v", err)

		os.Exit(-1)
	}

	cfg, isok := config.GetConfig()
	if !isok {
		fmt.Printf("GetConfig fail!")

		os.Exit(-1)
	}

	s := StartServ(cfg.Service.Host)
	s.Wait()
}
