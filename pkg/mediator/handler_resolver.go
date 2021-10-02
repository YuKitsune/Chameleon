package mediator

import "reflect"

type HandlerPredicate = func (p reflect.Type) bool

type HandlerResolver interface {
	ResolveType(reflect.Type) (interface{}, error)
	ResolveMatchingType(HandlerPredicate) (interface{}, error)
	ResolveMatchingTypes(HandlerPredicate) ([]interface{}, error)
}

func FindHandlerForRequest(reqType reflect.Type/*, resType reflect.Type*/) HandlerPredicate {
	return func (t reflect.Type) bool {

		// If t has a "Handle" method that:
		// - Accepts reqType as it's only argument
		// - Returns only resType, or resType and error

		m, found := t.MethodByName(InternalHandlerFuncName)
		if !found {
			return false
		}

		numIn := m.Type.NumIn()
		if numIn != 1 {
			return false
		}

		foundReqType := m.Type.In(0)
		// foundResType := m.Type.Out(0)
		returnsError := m.Type.Out(1).Name() == "error"
		if !returnsError {
			return false
		}

		return foundReqType == reqType// && foundResType == resType
	}
}

func FindHandlerForEvent(evtType reflect.Type) HandlerPredicate {
	return func (t reflect.Type) bool {
		// If t has a "Handle" method that:
		// - Accepts evtType as it's only argument
		// - Returns nothing

		m, found := t.MethodByName(InternalHandlerFuncName)
		if !found {
			return false
		}

		numIn := m.Type.NumIn()
		if numIn != 1 {
			return false
		}

		foundEvtType := m.Type.In(0)
		returnsError := m.Type.Out(0).Name() == "error"
		if !returnsError {
			return false
		}

		return foundEvtType == evtType
	}
}