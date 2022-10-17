package di

import (
	some_contracts "github.com/ihatiko/di/test-data/some-contracts"
	some_data "github.com/ihatiko/di/test-data/some-data"
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
func TestProvide2Cyclomatic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.ErrorIs(t, r.(error), cyclomaticError)
		}
	}()
	ProvideInterface[some_data.Test5](some_data.NewConcreteTest5)
	ProvideInterface[some_data.Test6](some_data.NewConcreteTest6)
	Invoke(func(t5 some_data.Test5, t6 some_data.Test6) {

	})
}

func TestProvide2Deep(t *testing.T) {
	ProvideInterface[some_data.Test1](some_data.NewConcreteTest1)
	ProvideInterface[some_data.Test2](some_data.NewConcreteTest2)
	ProvideInterface[some_data.Test3](some_data.NewConcreteTest3)
	ProvideInterface[some_data.Test4](some_data.NewConcreteTest4)

	Invoke(func(t1 some_data.Test1, t2 some_data.Test2, t3 some_data.Test3, t4 some_data.Test4) {
		assert.NotEqual(t, t1, nil)
		assert.NotEqual(t, t2, nil)
		assert.NotEqual(t, t3, nil)
		assert.NotEqual(t, t4, nil)
	})
}
