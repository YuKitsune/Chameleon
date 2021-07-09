package errors

type EntityNotFoundError struct {
	Params interface{}
}

func NewEntityNotFoundError(params interface{}) *EntityNotFoundError {
	return &EntityNotFoundError{params}
}

func (err *EntityNotFoundError) Error() string {
	return "no matching record could be found"
}