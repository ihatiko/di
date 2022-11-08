# DI
A reflection generic based [Go](http://golang.org) [Dependency injection](https://en.wikipedia.org/wiki/Dependency_injection).
### Require
Golang [1.18+](https://go.dev/blog/go1.18)

### Reason for using
* Simplify code
* Nothing string keys 
* Automatic create objects with dependencies
* Lazy injections
* You can use Concrete objects or Function Constructor for inject

### Good for:

* Resolving the object graph during process startup.

### Bad for:

* Resolving dependencies after the process has already started.


## Installation

```bash
go get github.com/ihatiko/di
```

## Basic Usage
### Test Data 
```go

type Service interface {
	ServiceTest()
}

type Handler interface {
	HandlerTest()
}

type ConcreteHandler struct {
	Service some_contracts.Service
}

func NewConcreteHandler(service some_contracts.Service) *ConcreteHandler {
	return &ConcreteHandler{Service: service}
}

func (h ConcreteHandler) HandlerTest() {
	fmt.Println("Hello world HandlerTest")
}

type ConcreteService struct {
	Repository some_contracts.Repository
}

func NewConcreteService(repository some_contracts.Repository) *ConcreteService {
	return &ConcreteService{Repository: repository}
}

func (s ConcreteService) ServiceTest() {
	fmt.Println("Hello world ServiceTest")
}
```
### Provide interface
```go
import "get github.com/ihatiko/di"

di.ProvideInterface[some_contracts.Service](some_injection.ConcreteService{})
di.ProvideInterface[some_contracts.Handler](some_injection.NewConcreteHandler)

di.Invoke(func(h some_contracts.Handler, s some_contracts.Service) {
    h.HandlerTest()
    s.ServiceTest()
})
```

```go
import "get github.com/ihatiko/di"

cfg :=
di.Provide(
    &some_injection.ConfigRepository{},
    some_injection.NewConcreteRepository,
)
data := GetInject[*some_injection.ConcreteRepository]()
```

#### [Some examples in tests](https://github.com/ihatiko/di/blob/main/di_test.go)
