package mediator

import (
	"fmt"
	"reflect"
)

const (
	InternalHandlerFuncName = "Handle"
)

type ErrHandlerMethodNotFound struct {
}

func (err *ErrHandlerMethodNotFound) Error() string {
	return "could not find handler method"
}

type handler interface {
	Invoke(r interface{}) error
	GetHandlerType() reflect.Type
	GetRequestType() reflect.Type
}

func getHandlerMethodAndRequestType(v interface{}) (*reflect.Value, *reflect.Type, error) {

	fmt.Printf("%T", v)
	method := reflect.ValueOf(v).MethodByName(InternalHandlerFuncName)
	numIn := method.Type().NumIn()
	if numIn != 1 {
		return nil, nil, &ErrHandlerMethodNotFound{}
	}

	requestType := method.Type().In(0)
	returnsError := method.Type().Out(0).Name() == "error"
	if !returnsError {
		return nil, nil, &ErrHandlerMethodNotFound{}
	}

	return &method, &requestType, nil
}

func getRequestType(handlerType reflect.Type) (*reflect.Type, error) {

	method, exists := handlerType.MethodByName(InternalHandlerFuncName)
	if !exists {
		return nil, &ErrHandlerMethodNotFound{}
	}

	numIn := method.Type.NumIn()
	if numIn != 2 {
		return nil, &ErrHandlerMethodNotFound{}
	}

	// first arg is the received, second is the event
	requestType := method.Type.In(1)
	returnsError := method.Type.Out(0).Name() == "error"
	if !returnsError {
		return nil, &ErrHandlerMethodNotFound{}
	}

	return &requestType, nil
}