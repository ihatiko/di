package some_injection

import (
	"fmt"
	some_contracts "github.com/ihatiko/di/test-data/some-contracts"
)

type ConcreteHandler struct {
	Service some_contracts.Service
}

func NewConcreteHandler(service some_contracts.Service) *ConcreteHandler {
	return &ConcreteHandler{Service: service}
}

func (h ConcreteHandler) HandlerTest(test string) {
	fmt.Println("Hello world HandlerTest")
}
