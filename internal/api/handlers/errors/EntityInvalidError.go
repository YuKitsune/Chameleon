package errors

import (
	"fmt"
)

type EntityInvalidError struct {
	Entity interface{}
	inner  error
}

func NewEntityInvalidErrorFromErr(entity interface{}, err error) *EntityInvalidError {
	return &EntityInvalidError{entity, err}
}

func NewEntityInvalidError(entity interface{}, format string, args ...interface{}) *EntityInvalidError {
	return &EntityInvalidError{entity, fmt.Errorf(format, args)}
}

func (err *EntityInvalidError) Error() string {
	return "the entity is invalid: %v"
}
