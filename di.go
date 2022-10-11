package di

import (
	"fmt"
	"reflect"
)

type Registry struct {
	Injections map[string]reflect.Value
}

var registry = &Registry{Injections: map[string]reflect.Value{}}

func Provide(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorValueType := reflect.ValueOf(constructor)
		constructorType := reflect.TypeOf(constructor)
		pkgPath := constructorType.PkgPath()
		if constructorType.Kind() == reflect.Pointer {
			pkgPath = constructorType.Elem().PkgPath()
		}
		key := ""
		if constructorType.Kind() == reflect.Func {
			outType := constructorType.Out(0)
			if outType.Kind() == reflect.Pointer {
				pkgPath = outType.Elem().PkgPath()
			}
			key = fmt.Sprintf("%s_%s", pkgPath, outType.String())
		} else {
			key = fmt.Sprintf("%s_%s", pkgPath, constructorType.String())
		}

		baseInject(constructorType, constructorValueType, key)
	}
}

func GetInject[T any]() T {
	typeA := reflect.TypeOf((*T)(nil)).Elem()
	pkgPath := typeA.PkgPath()
	if typeA.Kind() == reflect.Pointer {
		pkgPath = typeA.Elem().PkgPath()
	}
	key := fmt.Sprintf("%s_%s", pkgPath, typeA.String())
	if data, ok := registry.Injections[key]; ok {
		return data.Interface().(T)
	} else {
		panic(fmt.Sprintf("Injection with key %s does not exist", key))
	}
}

func ProvideInterface[T any](constructor any) {
	constructorValueType := reflect.ValueOf(constructor)
	constructorType := reflect.TypeOf(constructor)
	typeA := reflect.TypeOf((*T)(nil)).Elem()
	key := fmt.Sprintf("%s_%s", typeA.PkgPath(), typeA.String())

	if typeA.Kind() != reflect.Interface {
		panic("Type T does not interface")
	}
	if constructorType.Kind() == reflect.Func {
		typeB := constructorValueType.Type().Out(0).Elem()
		if !typeB.Implements(typeA) {
			panic(fmt.Sprintf("Type A doesnt impletemnt type B %s %s", typeA.String(), typeB.String()))
		}
	}
	baseInject(constructorType, constructorValueType, key)
}

func baseInject(constructorType reflect.Type, constructorValueType reflect.Value, key string) {
	var injection reflect.Value
	if constructorType.Kind() == reflect.Func {
		var buffer []reflect.Value
		for i := 0; i < constructorType.NumIn(); i++ {
			inParam := constructorType.In(i)
			pkgPath := inParam.PkgPath()
			if inParam.Kind() == reflect.Pointer {
				pkgPath = inParam.Elem().PkgPath()
			}
			paramKey := fmt.Sprintf("%s_%s", pkgPath, constructorType.In(i).String())
			if inj, ok := registry.Injections[paramKey]; ok {
				buffer = append(buffer, inj)
			} else {
				panic(fmt.Sprintf("Does not exist %s dependencies", paramKey))
			}
		}
		data := constructorValueType.Call(buffer)
		if len(data) == 0 {
			panic("Constructor does not return injection")
		}
		injection = data[0]
	} else {
		injection = constructorValueType
	}
	registry.Injections[key] = injection
}

func Invoke(fns ...interface{}) {
	for _, fn := range fns {
		constructorValueType := reflect.ValueOf(fn)
		constructorType := reflect.TypeOf(fn)
		var buffer []reflect.Value
		for i := 0; i < constructorType.NumIn(); i++ {
			fnParam := reflect.TypeOf(fn).In(i)
			key := fmt.Sprintf("%s_%s", fnParam.PkgPath(), fnParam.String())
			if injection, ok := registry.Injections[key]; ok {
				buffer = append(buffer, injection)
			} else {
				panic(fmt.Sprintf("Injection with key %s does not exists", key))
			}
		}
		constructorValueType.Call(buffer)
	}
}
