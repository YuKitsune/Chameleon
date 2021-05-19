package mediator

import (
	"errors"
	"fmt"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"reflect"
)

type factoryHandler struct {
	container ioc.Container

	handlerType reflect.Type
	requestType reflect.Type

	factory interface{}
}

func newFactoryHandler(container ioc.Container, factory interface{}) (*factoryHandler, error) {
	factoryType := reflect.TypeOf(factory)
	if factoryType == nil {
		return nil, errors.New("can't provide an untyped nil")
	}
	if factoryType.Kind() != reflect.Func {
		return nil, errors.New(fmt.Sprintf("must provide factory function, got %v (type %v)", factory, factoryType))
	}

	handlerType := reflect.ValueOf(factory).Type().Out(0)
	requestType, err := getRequestType(handlerType)
	if err != nil {
		return nil, err
	}

	// Todo: I'm not overly happy that we're mutating the container, but we need it's dependencies...
	err = container.RegisterTransientFactory(factory)
	if err != nil {
		return nil, err
	}

	return &factoryHandler{
		container:   container,
		handlerType: handlerType,
		requestType: *requestType,
		factory:     factory,
	}, nil
}

func (h *factoryHandler) Invoke(r interface{}) error {

	// Make the function type that we pass into our container to resolve the handler
	// First type is the receiver, the rest are the args
	fnIn := []reflect.Type {h.handlerType}
	fnOut := []reflect.Type {reflect.ValueOf(errors.New).Type().Out(0)}
	fnType := reflect.FuncOf(fnIn, fnOut, false)

	// Make the function
	fn := reflect.MakeFunc(
		fnType,
		func(args []reflect.Value) (results []reflect.Value) {

			// Get the handlers instance
			handlerInstance := args[0]
			method, requestType, err := getHandlerMethodAndRequestType(handlerInstance.Interface())
			if err != nil {
				return []reflect.Value {reflect.ValueOf(err)}
			}

			// Ensure the handler we got can handle the request we've received
			if *requestType != reflect.TypeOf(r) {
				return []reflect.Value {reflect.ValueOf(ErrHandlerMethodNotFound{})}
			}

			in := []reflect.Value { reflect.ValueOf(r) }
			return method.Call(in)
		})

	err := h.container.ResolveInScope(fn.Interface())
	if err != nil {
		return err
	}

	return nil
}

func (h *factoryHandler) GetRequestType() reflect.Type {
	return h.requestType
}

func (h *factoryHandler) GetHandlerType() reflect.Type {
	return h.handlerType
}
