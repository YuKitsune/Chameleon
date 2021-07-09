package errors

type EntityExistsError struct {
	Entity interface{}
}

func NewEntityExistsError(entity interface{}) *EntityExistsError {
	return &EntityExistsError{entity}
}

func (err *EntityExistsError) Error() string {
	return "the entity already exists"
}