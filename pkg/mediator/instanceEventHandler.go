package mediator

import (
"reflect"
)


type instanceEventHandler struct {
	instance       interface{}
	eventType   reflect.Type
	handlerMethod reflect.Value
}

func newInstanceEventHandler(h interface{}) (*instanceEventHandler, error) {

	method, requestType, err := getHandlerMethodAndEventType(h)
	if err != nil {
		return nil, err
	}

	return &instanceEventHandler{
		instance:       h,
		eventType:   *requestType,
		handlerMethod: *method,
	}, nil
}

func (h *instanceEventHandler) Invoke(r interface{}) error {

	// Setup our argument
	in := []reflect.Value {reflect.ValueOf(r)}

	// Invoke the handler
	out := h.handlerMethod.Call(in)
	if !out[0].IsNil() {
		return out[0].Interface().(error)
	}

	return nil
}

func (h *instanceEventHandler) GetEventType() reflect.Type {
	return h.eventType
}

func (h *instanceEventHandler) GetHandlerType() reflect.Type {
	return reflect.TypeOf(h.instance)
}
