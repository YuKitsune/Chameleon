package model

import (
	"fmt"
	"reflect"
)

type CastFailedError struct {
	targetType reflect.Type
}

func CastFailedErr(t reflect.Type) *CastFailedError {
	return &CastFailedError{targetType: t}
}

func (e *CastFailedError) Error() string {
	return fmt.Sprintf("failed to cast to type %s", e.targetType.Name())
}
