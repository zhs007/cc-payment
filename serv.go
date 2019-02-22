package main

import (
	"context"
	"net/http"

	"github.com/zhs007/cc-payment/logger"
	"github.com/zhs007/cc-payment/router"
	"go.uber.org/zap"
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
			// ReadTimeout:    10 * time.Second,
			// WriteTimeout:   10 * time.Second,
			// MaxHeaderBytes: 1 << 20,
		},
		done: make(chan int),
	}

	go func() {
		err := s.serv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("ListenAndServe error.", zap.Error(err))
		}

		// fmt.Print("StartServ done\n")
		// s.done <- 0
	}()

	return s
}

// Wait - wait service done
func (s *Serv) Wait() {
	<-s.done
}

// Stop - stop service
func (s *Serv) Stop() {
	// s.serv.Close()
	err := s.serv.Shutdown(context.Background())
	if err != nil {
		logger.Error("Serv.Stop:Shutdown err %v", zap.Error(err))
	}

	s.done <- 0
}
