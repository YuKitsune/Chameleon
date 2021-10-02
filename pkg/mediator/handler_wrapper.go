package mediator

import "reflect"

const (
	InternalHandlerFuncName = "Handle"
)

type requestHandlerWrapper struct {
	method reflect.Value
	handlerType reflect.Type
}

func newRequestHandlerWrapper(handler interface{}) *requestHandlerWrapper {
	val := reflect.ValueOf(handler)
	typ := val.Type()
	method := val.MethodByName(InternalHandlerFuncName)
	return &requestHandlerWrapper{
		method,
		typ,
	}
}

func (w *requestHandlerWrapper) InvokeAsRequestHandler(req interface{}) (interface{}, error) {
	in := []reflect.Value{reflect.ValueOf(req)}
	out := w.method.Call(in)

	res := valueOrNil(out[0])
	err := errorOrNil(out[1])

	return res, err
}

func (w *requestHandlerWrapper) InvokeAsEventHandler(req interface{}) error {
	in := []reflect.Value{reflect.ValueOf(req)}
	out := w.method.Call(in)

	err := errorOrNil(out[0])

	return err
}

func (w *requestHandlerWrapper) HandlerType() reflect.Type {
	return w.handlerType
}
