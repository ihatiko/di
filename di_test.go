package di

import (
	some_contracts "github.com/ihatiko/di/test-data/some-contracts"
	some_injection "github.com/ihatiko/di/test-data/some-injection"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInvoke__ProvideInterface(t *testing.T) {
	ProvideInterface[some_contracts.Service](some_injection.ConcreteService{})
	ProvideInterface[some_contracts.Handler](some_injection.NewConcreteHandler)

	Invoke(func(h some_contracts.Handler, s some_contracts.Service) {
		h.HandlerTest()
		s.ServiceTest()
	})
}

func TestGetInject(t *testing.T) {
	s := &some_injection.ConcreteService{}
	ProvideInterface[some_contracts.Service](s)
	data := GetInject[some_contracts.Service]()
	assert.Equal(t, reflect.TypeOf(s).String(), reflect.TypeOf(data).String())
}

func TestProvide(t *testing.T) {
	cfg := &some_injection.ConfigRepository{}
	Provide(
		cfg,
		some_injection.NewConcreteRepository,
	)
	data := GetInject[*some_injection.ConcreteRepository]()
	assert.Equal(t, reflect.TypeOf(&some_injection.ConcreteRepository{}).String(), reflect.TypeOf(data).String())
}

func TestProvide2(t *testing.T) {
	cfg := &some_injection.ConfigRepository{}
	Provide(
		some_injection.NewConcreteRepository,
		cfg,
	)
	data := GetInject[*some_injection.ConcreteRepository]()
	assert.Equal(t, reflect.TypeOf(&some_injection.ConcreteRepository{}).String(), reflect.TypeOf(data).String())
}
