package mediator

import (
	"fmt"
	"github.com/yukitsune/chameleon/pkg/errors"
	"reflect"
)

type ErrHandlerNotFound struct {
	RequestType reflect.Type
}

func (err *ErrHandlerNotFound) Error() string {
	return fmt.Sprintf("No handler has been registered for request type %v", err.RequestType)
}

type Mediator struct {
	handlers []*handler
}

func New() *Mediator {
	return &Mediator{}
}

func (m *Mediator) AddHandler(v interface{}) error {
	h, err := newHandler(v)
	if err != nil {
		return err
	}

	m.handlers = append(m.handlers, h)
	return nil
}

func (m *Mediator) Send(r interface{}) error {

	handlers := m.findHandlers(r)
	var errs errors.Errors
	for _, h := range handlers {
		err := h.Invoke(r)
		if err != nil {
			wrappedErr := NewError(h.handler, err)
			errs = append(errs, wrappedErr)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (m *Mediator) findHandlers(r interface{}) []*handler {
	var handlers []*handler

	requestType := reflect.TypeOf(r)
	for _, h := range m.handlers {
		if h.requestType == requestType {
			handlers = append(handlers, h)
		}
	}

	return handlers
}
