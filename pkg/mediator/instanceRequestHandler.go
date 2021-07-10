package mediator

import (
	"reflect"
)

type instanceRequestHandler struct {
	instance      interface{}
	requestType   reflect.Type
	responseType  reflect.Type
	handlerMethod reflect.Value
}

func newInstanceRequestHandler(h interface{}) (*instanceRequestHandler, error) {

	method, requestType, responseType, err := getHandlerMethodAndRequestAndResponseType(h)
	if err != nil {
		return nil, err
	}

	return &instanceRequestHandler{
		instance:      h,
		requestType:   *requestType,
		responseType:  *responseType,
		handlerMethod: *method,
	}, nil
}

func (h *instanceRequestHandler) Invoke(r interface{}) (interface{}, error) {

	// Setup our argument
	in := []reflect.Value{reflect.ValueOf(r)}

	// Invoke the handler
	out := h.handlerMethod.Call(in)
	if !out[1].IsNil() {
		return nil, out[1].Interface().(error)
	}

	return out[0].Interface(), nil
}

func (h *instanceRequestHandler) GetRequestType() reflect.Type {
	return h.requestType
}
func (h *instanceRequestHandler) GetResponseType() reflect.Type {
	return h.responseType
}

func (h *instanceRequestHandler) GetHandlerType() reflect.Type {
	return reflect.TypeOf(h.instance)
}
