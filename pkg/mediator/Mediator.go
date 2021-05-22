package mediator

import (
	"fmt"
	"github.com/yukitsune/chameleon/pkg/errors"
	"github.com/yukitsune/chameleon/pkg/ioc"
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
	container ioc.Container
}

func New(c ioc.Container) *Mediator {
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
