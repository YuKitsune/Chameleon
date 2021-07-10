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

type eventHandler interface {
	Invoke(r interface{}) error
	GetHandlerType() reflect.Type
	GetEventType() reflect.Type
}

func getHandlerMethodAndEventType(v interface{}) (*reflect.Value, *reflect.Type, error) {
	method := reflect.ValueOf(v).MethodByName(InternalHandlerFuncName)
	numIn := method.Type().NumIn()
	if numIn != 1 {
		return nil, nil, &ErrHandlerMethodNotFound{}
	}

	eventType := method.Type().In(0)
	returnsError := method.Type().Out(0).Name() == "error"
	if !returnsError {
		return nil, nil, &ErrHandlerMethodNotFound{}
	}

	return &method, &eventType, nil
}

func getEventOrRequestType(handlerType reflect.Type) (*reflect.Type, error) {

	method, exists := handlerType.MethodByName(InternalHandlerFuncName)
	if !exists {
		return nil, &ErrHandlerMethodNotFound{}
	}

	numIn := method.Type.NumIn()
	if numIn != 2 {
		return nil, &ErrHandlerMethodNotFound{}
	}

	// first arg is the receiver, second is the event
	eventType := method.Type.In(1)
	return &eventType, nil
}
