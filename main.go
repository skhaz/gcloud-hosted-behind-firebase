package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
)

type ArithService struct{}

func (s *ArithService) Add(a, b float64) float64 {
	return a + b
}

func (s *ArithService) Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func (s *ArithService) Version() (string, error) {
	return fmt.Sprintf("%d", 1), nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var server = rpc.NewServer()

	if err := server.RegisterName("arith", &ArithService{}); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server)
}
