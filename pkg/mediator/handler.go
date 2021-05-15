package mediator

import (
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

type handler struct {
	handler       interface{}
	requestType   reflect.Type
	handlerMethod reflect.Value
}

func newHandler(h interface{}) (*handler, error) {

	method := reflect.ValueOf(h).MethodByName(InternalHandlerFuncName)
	numIn := method.Type().NumIn()
	if numIn != 1 {
		return nil, &ErrHandlerMethodNotFound{}
	}

	requestType := method.Type().In(0)
	returnsError := method.Type().Out(0).Name() == "error"
	if !returnsError {
		return nil, &ErrHandlerMethodNotFound{}
	}

	return &handler{
		handler:       h,
		requestType:   requestType,
		handlerMethod: method,
	}, nil
}

func (h *handler) Invoke(r interface{}) error {

	// Setup our argument
	var in []reflect.Value
	in = append(in, reflect.ValueOf(r))

	// Invoke the handler
	out := h.handlerMethod.Call(in)
	if !out[0].IsNil() {
		return out[0].Interface().(error)
	}

	return nil
}
