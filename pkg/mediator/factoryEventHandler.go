package mediator

import (
	"errors"
	"fmt"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"reflect"
)

type factoryEventHandler struct {
	container ioc.Container

	handlerType reflect.Type
	eventType reflect.Type

	factory interface{}
}

func newFactoryEventHandler(container ioc.Container, factory interface{}) (*factoryEventHandler, error) {
	factoryType := reflect.TypeOf(factory)
	if factoryType == nil {
		return nil, errors.New("can't provide an untyped nil")
	}
	if factoryType.Kind() != reflect.Func {
		return nil, errors.New(fmt.Sprintf("must provide factory function, got %v (type %v)", factory, factoryType))
	}

	handlerType := reflect.ValueOf(factory).Type().Out(0)
	eventType, err := getEventOrRequestType(handlerType)
	if err != nil {
		return nil, err
	}

	// Todo: I'm not overly happy that we're mutating the container, but we need it's dependencies...
	err = container.RegisterTransientFactory(factory)
	if err != nil {
		return nil, err
	}

	return &factoryEventHandler{
		container:   container,
		handlerType: handlerType,
		eventType: *eventType,
		factory:     factory,
	}, nil
}

func (h *factoryEventHandler) Invoke(r interface{}) error {

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
			method, eventType, err := getHandlerMethodAndEventType(handlerInstance.Interface())
			if err != nil {
				return []reflect.Value {reflect.ValueOf(err)}
			}

			// Ensure the handler we got can handle the request we've received
			if *eventType != reflect.TypeOf(r) {
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

func (h *factoryEventHandler) GetEventType() reflect.Type {
	return h.eventType
}

func (h *factoryEventHandler) GetHandlerType() reflect.Type {
	return h.handlerType
}
