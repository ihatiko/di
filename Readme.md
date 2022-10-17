# DI
A reflection generic based [Go](http://golang.org) [Dependency injection](https://en.wikipedia.org/wiki/Dependency_injection).
### Require
Golang [1.18+](https://go.dev/blog/go1.18)

### Reason for using
* Simplify code
* Nothing string keys 
* Automatic create objects with dependencies
* Lazy injections

### Good for:

* Resolving the object graph during process startup.

### Bad for:

* Resolving dependencies after the process has already started.


## Installation

```bash
go get github.com/ihatiko/di
```

## Basic Usage

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