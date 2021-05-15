package errors

type Errors []error

// implement the Error interface
func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	// multiple errors
	msg := ""
	for _, err := range e {
		msg += "\n" + err.Error()
	}
	return msg
}
