package some_injection

import "fmt"

type Service interface {
	ServiceTest()
}

type ConcreteService struct {
	Repository Repository
}

func NewConcreteService(repository Repository) *ConcreteService {
	return &ConcreteService{Repository: repository}
}

func (s ConcreteService) ServiceTest() {
	fmt.Println("Hello world ServiceTest")
}
