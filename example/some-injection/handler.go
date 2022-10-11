package some_injection

import "fmt"

type Handler interface {
	HandlerTest()
}

type ConcreteHandler struct {
	Service Service
}

func NewConcreteHandler(service Service) *ConcreteHandler {
	return &ConcreteHandler{Service: service}
}

func (h ConcreteHandler) HandlerTest() {
	fmt.Println("Hello world HandlerTest")
}
