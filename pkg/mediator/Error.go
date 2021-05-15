package mediator

import "fmt"

type Error struct {
	handler interface{}
	err     error
}

func NewError(handler interface{}, err error) *Error {
	return &Error{
		handler: handler,
		err:     err,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("Handler: %T Error: %s", e.handler, e.err.Error())
}
