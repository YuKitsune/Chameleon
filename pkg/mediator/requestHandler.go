package mediator

import (
	"fmt"
	"reflect"
)

type ErrRequestHandlerAlreadyRegistered struct {
	rt reflect.Type
}

func AlreadyRegistered(rt reflect.Type) *ErrRequestHandlerAlreadyRegistered {
	return &ErrRequestHandlerAlreadyRegistered{rt: rt}
}

func (err *ErrRequestHandlerAlreadyRegistered) Error() string {
	return fmt.Sprintf("a request handler of type %s has already been registered", err.rt.Name())
}

type requestHandler interface {
	Invoke(r interface{}) (interface{}, error)
	GetHandlerType() reflect.Type
	GetRequestType() reflect.Type
	GetResponseType() reflect.Type
}

func getHandlerMethodAndRequestAndResponseType(v interface{}) (*reflect.Value, *reflect.Type, *reflect.Type, error) {

	method := reflect.ValueOf(v).MethodByName(InternalHandlerFuncName)
	numIn := method.Type().NumIn()
	if numIn != 1 {
		return nil, nil, nil, &ErrHandlerMethodNotFound{}
	}

	requestType := method.Type().In(0)
	responseType := method.Type().Out(0)
	returnsError := method.Type().Out(1).Name() == "error"
	if !returnsError {
		return nil, nil, nil, &ErrHandlerMethodNotFound{}
	}

	return &method, &requestType, &responseType, nil
}

func getResponseType(handlerType reflect.Type) (*reflect.Type, error) {

	method, exists := handlerType.MethodByName(InternalHandlerFuncName)
	if !exists {
		return nil, &ErrHandlerMethodNotFound{}
	}

	// first arg is the receiver, second is the request
	numIn := method.Type.NumIn()
	if numIn != 2 {
		return nil, &ErrHandlerMethodNotFound{}
	}

	responseType := method.Type.Out(0)
	returnsError := method.Type.Out(1).Name() == "error"
	if !returnsError {
		return nil, &ErrHandlerMethodNotFound{}
	}

	return &responseType, nil
}