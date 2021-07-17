package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MissingParameterError struct {
	parameterName string
}

func MissingParameterErr(parameterName string) *MissingParameterError {
	return &MissingParameterError{
		parameterName: parameterName,
	}
}

func (e *MissingParameterError) Error() string {
	return fmt.Sprintf("could not find parameter \"%s\"", e.parameterName)
}

func getRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
