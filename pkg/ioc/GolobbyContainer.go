package ioc

import (
	"github.com/golobby/container/pkg/container"
	"reflect"
)

// Todo: Looking for the c.Call method that returns an error but it doesn't exist...

type GolobbyContainer struct {
	c *container.Container
}

func NewGolobbyContainer() *GolobbyContainer {
	c := container.NewContainer()
	return &GolobbyContainer{c: &c}
}

func (g GolobbyContainer) RegisterSingletonInstance(v interface{}) error {
	resolver := makeFuncThatReturns(v)
	g.c.Singleton(resolver)
	return nil
}

func (g GolobbyContainer) RegisterSingletonFactory(v interface{}) error {
	g.c.Singleton(v)
	return nil
}

func (g GolobbyContainer) RegisterTransientInstance(v interface{}) error {
	resolver := makeFuncThatReturns(v)
	g.c.Singleton(resolver)
	return nil
}

func (g GolobbyContainer) RegisterTransientFactory(v interface{}) error {
	g.c.Transient(v)
	return nil
}

func (g GolobbyContainer) ResolveInScope(v interface{}) error {
	var err error
	errAddr := &err
	g.c.Make(wrapFunc(v, errAddr))
	return err
}

func makeFuncThatReturns(v interface{}) interface{} {
	instanceType := reflect.TypeOf(v)
	fnType := reflect.FuncOf([]reflect.Type { }, []reflect.Type { instanceType }, false)
	fn := reflect.MakeFunc(
		fnType,
		func (args []reflect.Value) []reflect.Value {
			return []reflect.Value { reflect.ValueOf(v) }
		})

	return fn.Interface()
}


func wrapFunc(v interface{}, err *error) interface{} {
	originalFnType := reflect.TypeOf(v)
	originalFnValue := reflect.ValueOf(v)

	var in []reflect.Type
	for i := 0; i < originalFnType.NumIn(); i++ {
		in = append(in, originalFnType.In(i))
	}

	var out []reflect.Type
	for i := 0; i < originalFnType.NumOut(); i++ {
		out = append(out, originalFnType.Out(i))
	}

	fnType := reflect.FuncOf(in, out, false)
	fn := reflect.MakeFunc(
		fnType,
		func (args []reflect.Value) []reflect.Value {

			out := originalFnValue.Call(args)

			// If we have an error, then assign it
			if len(out) > 0 {
				maybeError := out[0]
				if !maybeError.IsNil() {
					if actualError, ok := maybeError.Interface().(error); ok {
						*err = actualError
					}
				}
			}

			return out
		})

	return fn.Interface()
}