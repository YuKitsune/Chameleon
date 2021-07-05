package mediator

import (
	"fmt"
	"github.com/yukitsune/camogo"
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
	eventHandlers []eventHandler
	requestHandlers []requestHandler
	container camogo.Container
}

func New(c camogo.Container) *Mediator {
	return &Mediator{
		container: c,
	}
}

func (m *Mediator) AddEventHandlerInstance(v interface{}) error {
	h, err := newInstanceEventHandler(v)
	if err != nil {
		return err
	}

	m.eventHandlers = append(m.eventHandlers, h)
	return nil
}

func (m *Mediator) AddEventHandlerFactory(ctor interface{}) error {
	h, err := newFactoryEventHandler(m.container, ctor)
	if err != nil {
		return err
	}

	m.eventHandlers = append(m.eventHandlers, h)
	return nil
}

func (m *Mediator) AddRequestHandlerInstance(v interface{}) error {
	h, err := newInstanceRequestHandler(v)
	if err != nil {
		return err
	}

	if m.requestHandlerExists(h.GetRequestType()) {
		return AlreadyRegistered(h.GetRequestType())
	}

	m.requestHandlers = append(m.requestHandlers, h)
	return nil
}

func (m *Mediator) AddRequestHandlerFactory(ctor interface{}) error {
	h, err := newFactoryRequestHandler(m.container, ctor)
	if err != nil {
		return err
	}

	if m.requestHandlerExists(h.GetRequestType()) {
		return AlreadyRegistered(h.GetRequestType())
	}

	m.requestHandlers = append(m.requestHandlers, h)
	return nil
}

func (m *Mediator) Send(r interface{}) (interface{}, error) {

	handler := m.findRequestHandler(r)

	if handler == nil {
		return nil, nil
	}

	res, err := handler.Invoke(r)
	if err != nil {
		wrappedErr := NewError(handler.GetHandlerType(), err)
		return nil, wrappedErr
	}

	return res, nil
}

func (m *Mediator) Publish(r interface{}) error {

	handlers := m.findEventHandlers(r)
	var errs errors.Errors
	for _, h := range handlers {
		err := h.Invoke(r)
		if err != nil {
			wrappedErr := NewError(h.GetHandlerType(), err)
			errs = append(errs, wrappedErr)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (m *Mediator) findEventHandlers(r interface{}) []eventHandler {
	var handlers []eventHandler

	requestType := reflect.TypeOf(r)
	for _, h := range m.eventHandlers {
		if h.GetEventType() == requestType {
			handlers = append(handlers, h)
		}
	}

	return handlers
}

func (m *Mediator) findRequestHandler(r interface{}) requestHandler {
	requestType := reflect.TypeOf(r)

	for _, h := range m.requestHandlers {
		if h.GetRequestType() == requestType {
			return h
		}
	}

	return nil
}

func (m *Mediator) requestHandlerExists(rt reflect.Type) bool {
	for _, handler := range m.requestHandlers {
		if handler.GetRequestType() == rt {
			return true
		}
	}

	return false
}