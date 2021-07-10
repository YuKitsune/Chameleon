package mediator

import (
	"fmt"
	"reflect"
)

type Error struct {
	handlerType reflect.Type
	err         error
}

func NewError(handlerType reflect.Type, err error) *Error {
	return &Error{
		handlerType: handlerType,
		err:         err,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("Handler: %s Error: %s", e.handlerType.String(), e.err.Error())
}
