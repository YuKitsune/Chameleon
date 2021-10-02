package mediator

import (
	"github.com/yukitsune/chameleon/pkg/errors"
	"reflect"
)

type Mediator interface {
	Send(r interface{}) (interface{}, error)
	Publish(r interface{}) error
}

type mediator struct {
	resolver HandlerResolver
}

func (m *mediator) Send(req interface{}) (interface{}, error) {

	handler, err := m.resolver.ResolveMatchingType(FindHandlerForRequest(reflect.TypeOf(req)))
	if err != nil {
		return nil, err
	}

	wrappedHandler := newRequestHandlerWrapper(handler)
	res, err := wrappedHandler.InvokeAsRequestHandler(req)
	if err != nil {
		return nil, NewError(wrappedHandler.HandlerType(), err)
	}

	return res, nil
}

func (m *mediator) Publish(evt interface{}) error {

	handlers, err := m.resolver.ResolveMatchingTypes(FindHandlerForEvent(reflect.TypeOf(evt)))
	if err != nil {
		return err
	}

	var errs errors.Errors
	for _, h := range handlers {
		wrappedHandler := newRequestHandlerWrapper(h)
		err := wrappedHandler.InvokeAsEventHandler(evt)
		if err != nil {
			wrappedErr := NewError(wrappedHandler.HandlerType(), err)
			errs = append(errs, wrappedErr)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
