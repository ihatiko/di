package di

import (
	some_injection "github.com/ihatiko/di/example/some-injection"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInvoke__ProvideInterface(t *testing.T) {
	ProvideInterface[some_injection.Service](some_injection.ConcreteService{})
	ProvideInterface[some_injection.Handler](some_injection.NewConcreteHandler)

	Invoke(func(h some_injection.Handler, s some_injection.Service) {
		h.HandlerTest()
		s.ServiceTest()
	})
}

func TestGetInject(t *testing.T) {
	s := &some_injection.ConcreteService{}
	ProvideInterface[some_injection.Service](s)
	data := GetInject[some_injection.Service]()
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
