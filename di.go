package di

import (
	"errors"
	"fmt"
	"reflect"
)

//TODO Scope, Transient, Singleton
type Registry struct {
	Scope        map[string]reflect.Value
	Dependencies map[string]Dependency
}

type Dependency struct {
	Constructor  interface{}
	Type         reflect.Type
	Dependencies []string
}

var registry = &Registry{
	Scope:        map[string]reflect.Value{},
	Dependencies: map[string]Dependency{},
}

func GetRegistry() *Registry {
	return registry
}

var cyclomaticError = errors.New("cyclomatic dependency")
var emptyInjectionError = errors.New("does not find injection")

func Clear() {
	registry = &Registry{
		Scope:        map[string]reflect.Value{},
		Dependencies: map[string]Dependency{},
	}
}

func Provide(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorValueType := reflect.ValueOf(constructor)
		constructorType := reflect.TypeOf(constructor)
		pkgPath := constructorType.PkgPath()
		if constructorType.Kind() == reflect.Pointer {
			pkgPath = constructorType.Elem().PkgPath()
		}
		inject(constructorType, pkgPath, constructor, constructorValueType)
	}
}

func inject(constructorType reflect.Type, pkgPath string, constructor interface{}, constructorValueType reflect.Value) {
	var inParams []string
	key := ""
	if constructorType.Kind() == reflect.Func {
		outType := constructorType.Out(0)
		if outType.Kind() == reflect.Pointer {
			pkgPath = outType.Elem().PkgPath()
		}
		for i := 0; i < constructorType.NumIn(); i++ {
			inParam := constructorType.In(i)
			inParamKey := inParam.PkgPath()
			if inParam.Kind() == reflect.Pointer {
				inParamKey = inParam.Elem().PkgPath()
			}
			inParams = append(inParams, fmt.Sprintf("%s/%s", inParamKey, inParam.String()))
		}
		key = fmt.Sprintf("%s/%s", pkgPath, outType.String())
		registry.Dependencies[key] = Dependency{
			Type:         constructorType,
			Dependencies: inParams,
			Constructor:  constructor,
		}
	} else {
		key = fmt.Sprintf("%s/%s", pkgPath, constructorType.String())
		registry.Scope[key] = constructorValueType
	}
}

func GetInject[T any]() T {
	typeA := reflect.TypeOf((*T)(nil)).Elem()
	pkgPath := typeA.PkgPath()
	if typeA.Kind() == reflect.Pointer {
		pkgPath = typeA.Elem().PkgPath()
	}
	key := fmt.Sprintf("%s/%s", pkgPath, typeA.String())
	if data, ok := registry.Scope[key]; ok {
		return data.Interface().(T)
	}
	data, err := buildInject(key)
	if errors.Is(err, cyclomaticError) {
		panic(err)
	}
	if errors.Is(err, emptyInjectionError) {
		panic(fmt.Sprintf("Injection with key %s does not exist", key))
	}
	return data.Interface().(T)
}

func findKey(key string, data []string) bool {
	result := false
	for _, item := range data {
		if item == key {
			return true
		}
	}
	return result
}
func buildInject(key string, paths ...string) (reflect.Value, error) {
	if findKey(key, paths) {
		return reflect.ValueOf(nil), cyclomaticError
	}
	if data, ok := registry.Dependencies[key]; ok {
		var injectionParams []reflect.Value
		for _, dpd := range data.Dependencies {
			if internalData, scopeOk := registry.Scope[dpd]; scopeOk {
				injectionParams = append(injectionParams, internalData)
				continue
			}
			paths = append(paths, key)
			inject, err := buildInject(dpd, paths...)
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			injectionParams = append(injectionParams, inject)
		}
		valueType := reflect.ValueOf(data.Constructor)
		result := valueType.Call(injectionParams)
		if len(result) == 0 {
			panic("Constructor does not return injection")
		}
		return result[0], nil
	} else {
		return reflect.ValueOf(nil), emptyInjectionError
	}
}

func ProvideInterface[T any](constructor any) {
	constructorValueType := reflect.ValueOf(constructor)
	constructorType := reflect.TypeOf(constructor)
	typeA := reflect.TypeOf((*T)(nil)).Elem()

	if typeA.Kind() != reflect.Interface {
		panic("Type T does not interface")
	}
	key := fmt.Sprintf("%s/%s", typeA.PkgPath(), typeA.String())
	if constructorType.Kind() == reflect.Func {
		typeB := constructorValueType.Type().Out(0)
		if typeB.Kind() == reflect.Pointer {
			typeB = typeB.Elem()
		}
		//TODO this function does not work with 'func(t *Test) Test()' where Test is pointer
		if !typeB.Implements(typeA) {
			panic(fmt.Sprintf("Type A doesnt implement type B %s %s", typeA.String(), typeB.String()))
		}
	}
	var inParams []string
	if constructorType.Kind() == reflect.Func {
		for i := 0; i < constructorType.NumIn(); i++ {
			inParam := constructorType.In(i)
			inParamKey := inParam.PkgPath()
			if inParam.Kind() == reflect.Pointer {
				inParamKey = inParam.Elem().PkgPath()
			}
			inParams = append(inParams, fmt.Sprintf("%s/%s", inParamKey, inParam.String()))
		}
		registry.Dependencies[key] = Dependency{
			Type:         constructorType,
			Dependencies: inParams,
			Constructor:  constructor,
		}
	} else {
		registry.Scope[key] = constructorValueType
	}
}

func Invoke(fns ...interface{}) {
	for _, fn := range fns {
		constructorValueType := reflect.ValueOf(fn)
		constructorType := reflect.TypeOf(fn)
		var buffer []reflect.Value
		for i := 0; i < constructorType.NumIn(); i++ {
			fnParam := reflect.TypeOf(fn).In(i)
			key := fmt.Sprintf("%s/%s", fnParam.PkgPath(), fnParam.String())
			if inj, ok := registry.Scope[key]; ok {
				buffer = append(buffer, inj)
				continue
			}
			data, err := buildInject(key)
			if errors.Is(err, cyclomaticError) {
				panic(err)
			}
			if errors.Is(err, emptyInjectionError) {
				panic(fmt.Sprintf("Injection with key %s does not exist", key))
			}
			buffer = append(buffer, data)
		}
		constructorValueType.Call(buffer)
	}
}
