package some_injection

import (
	"fmt"
	some_contracts "github.com/ihatiko/di/test-data/some-contracts"
)

type ConcreteService struct {
	Repository some_contracts.Repository
}

func NewConcreteService(repository some_contracts.Repository) *ConcreteService {
	return &ConcreteService{Repository: repository}
}

func NewAbstractService() some_contracts.Service {
	return &ConcreteService{}
}

func (s ConcreteService) ServiceTest() {
	fmt.Println("Hello world ServiceTest")
}
