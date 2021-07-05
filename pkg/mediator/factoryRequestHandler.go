package mediator

import (
	"errors"
	"fmt"
	"github.com/yukitsune/camogo"
	"reflect"
)

type factoryRequestHandler struct {
	container camogo.Container

	handlerType reflect.Type
	requestType reflect.Type
	responseType reflect.Type

	factory interface{}
}

func newFactoryRequestHandler(container camogo.Container, factory interface{}) (*factoryRequestHandler, error) {
	factoryType := reflect.TypeOf(factory)
	if factoryType == nil {
		return nil, errors.New("can't provide an untyped nil")
	}
	if factoryType.Kind() != reflect.Func {
		return nil, errors.New(fmt.Sprintf("must provide factory function, got %v (type %v)", factory, factoryType))
	}

	handlerType := reflect.ValueOf(factory).Type().Out(0)
	requestType, err := getEventOrRequestType(handlerType)
	if err != nil {
		return nil, err
	}

	responseType, err := getResponseType(handlerType)
	if err != nil {
		return nil, err
	}

	// Todo: I'm not overly happy that we're mutating the container, but we need it's dependencies...
	err = container.Register(func (r *camogo.Registrar) error {
		return r.RegisterFactory(factory, camogo.TransientLifetime)
	})
	if err != nil {
		return nil, err
	}

	return &factoryRequestHandler{
		container:   container,
		handlerType: handlerType,
		requestType: *requestType,
		responseType: *responseType,
		factory:     factory,
	}, nil
}

func (h *factoryRequestHandler) Invoke(r interface{}) (interface{}, error) {

	// Make the function type that we pass into our container to resolve the handler
	// First type is the receiver, the rest are the args
	fnIn := []reflect.Type {h.handlerType}
	fnOut := []reflect.Type {h.responseType, reflect.ValueOf(errors.New).Type().Out(0)}
	fnType := reflect.FuncOf(fnIn, fnOut, false)

	// Make the function
	fn := reflect.MakeFunc(
		fnType,
		func(args []reflect.Value) (results []reflect.Value) {

			// Get the handlers instance
			handlerInstance := args[0]
			method, requestType, _, err := getHandlerMethodAndRequestAndResponseType(handlerInstance.Interface())
			if err != nil {
				return []reflect.Value {reflect.ValueOf(nil), reflect.ValueOf(err)}
			}

			// Ensure the handler we got can handle the request we've received
			if *requestType != reflect.TypeOf(r) {
				return []reflect.Value {reflect.ValueOf(nil), reflect.ValueOf(ErrHandlerMethodNotFound{})}
			}

			in := []reflect.Value { reflect.ValueOf(r) }
			return method.Call(in)
		})

	res, err := h.container.ResolveWithResult(fn.Interface())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *factoryRequestHandler) GetRequestType() reflect.Type {
	return h.requestType
}

func (h *factoryRequestHandler) GetResponseType() reflect.Type {
	return h.responseType
}

func (h *factoryRequestHandler) GetHandlerType() reflect.Type {
	return h.handlerType
}
