package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

func GetParameter(name string, r *http.Request) (*string, error) {

	matchingParameters, ok := r.URL.Query()[name]

	if !ok || len(matchingParameters) == 0 {
		err := errors.New(fmt.Sprintf("parameter '%s' missing", name))
		return nil, err
	}

	if len(matchingParameters) > 1 {
		err := errors.New(fmt.Sprintf("multiple '%s' parameters found, expected 1", name))
		return nil, err
	}

	parameter := matchingParameters[0]

	return &parameter, nil
}
