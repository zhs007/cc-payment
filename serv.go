package main

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/zhs007/cc-payment/logger"
	"github.com/zhs007/cc-payment/router"
)

// Serv - service
type Serv struct {
	serv *http.Server
	done chan int
}

// StartServ -
func StartServ(servaddr string) *Serv {
	router.SetRouter()

	s := &Serv{
		serv: &http.Server{
			Addr:    servaddr,
			Handler: router.GinEngine,
		},
		done: make(chan int),
	}

	go func() {
		err := s.serv.ListenAndServe()
		if err != nil {
			logger.Error("ListenAndServe error.", zap.Error(err))
		}

		s.done <- 0
	}()

	return s
}

// Wait - wait service done
func (s *Serv) Wait() {
	<-s.done
}

// Stop - stop service
func (s *Serv) Stop() {
	s.serv.Shutdown(context.Background())
}
