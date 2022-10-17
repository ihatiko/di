package some_injection

import "fmt"

type ConfigRepository struct {
}

type ConcreteRepository struct {
	Cfg *ConfigRepository
}

func NewConcreteRepository(cfg *ConfigRepository) *ConcreteRepository {
	return &ConcreteRepository{Cfg: cfg}
}

func (s ConcreteRepository) RepositoryTest() {
	fmt.Println("Hello world Repository")
}
