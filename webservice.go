package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo"
)

type WebService struct {
	e       *echo.Echo
	context *MainContext
}

func NewWebService(c *MainContext) *WebService {
	ec := echo.New()
	// TODO init
	return &WebService{
		e:       ec,
		context: c,
	}
}

func (s *WebService) Start(port int, hostname string) {
	addr := fmt.Sprintf(hostname, ":", port)
	go func() {
		log.Println("server listens at http://", hostname, ":", port)
		if err := s.e.Start(addr); err != nil {
			log.Println("shutting down the server")
		}
	}()
}

func (s *WebService) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
