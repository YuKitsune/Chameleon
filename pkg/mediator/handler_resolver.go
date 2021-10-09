package mediator

import (
	"reflect"
)

type HandlerPredicate = func (p reflect.Type) bool

type HandlerResolver interface {
	ResolveType(reflect.Type) (interface{}, error)
	ResolveMatchingType(HandlerPredicate) (interface{}, error)
	ResolveMatchingTypes(HandlerPredicate) ([]interface{}, error)
}

func FindHandlerForRequest(reqType reflect.Type) HandlerPredicate {
	return func (t reflect.Type) bool {

		// If t has a "Handle" method that:
		// - Accepts reqType as it's only argument
		// - Returns a response and an error

		m, found := t.MethodByName(InternalHandlerFuncName)
		if !found {
			return false
		}

		// [0]: Receiver
		// [1]: Request Type
		numIn := m.Type.NumIn()
		if numIn != 2 {
			return false
		}

		foundReqType := m.Type.In(1)
		returnsError := m.Type.Out(1).Name() == "error"
		if !returnsError {
			return false
		}

		return foundReqType == reqType
	}
}

func FindHandlerForEvent(evtType reflect.Type) HandlerPredicate {
	return func (t reflect.Type) bool {
		// If t has a "Handle" method that:
		// - Accepts evtType as it's only argument
		// - Returns an error

		m, found := t.MethodByName(InternalHandlerFuncName)
		if !found {
			return false
		}

		// [0]: Receiver
		// [1]: Request Type
		numIn := m.Type.NumIn()
		if numIn != 2 {
			return false
		}

		foundEvtType := m.Type.In(1)
		returnsError := m.Type.Out(0).Name() == "error"
		if !returnsError {
			return false
		}

		return foundEvtType == evtType
	}
}