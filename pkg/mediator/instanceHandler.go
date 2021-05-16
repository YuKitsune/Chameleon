package mediator

import (
"reflect"
)


type instanceHandler struct {
	instance       interface{}
	requestType   reflect.Type
	handlerMethod reflect.Value
}

func newInstanceHandler(h interface{}) (*instanceHandler, error) {

	method, requestType, err := getHandlerMethodAndRequestType(h)
	if err != nil {
		return nil, err
	}

	return &instanceHandler{
		instance:       h,
		requestType:   *requestType,
		handlerMethod: *method,
	}, nil
}

func (h *instanceHandler) Invoke(r interface{}) error {

	// Setup our argument
	in := []reflect.Value {reflect.ValueOf(r)}

	// Invoke the handler
	out := h.handlerMethod.Call(in)
	if !out[0].IsNil() {
		return out[0].Interface().(error)
	}

	return nil
}

func (h *instanceHandler) GetRequestType() reflect.Type {
	return h.requestType
}

func (h *instanceHandler) GetHandlerType() reflect.Type {
	return reflect.TypeOf(h.instance)
}
