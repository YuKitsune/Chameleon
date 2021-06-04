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
	return &MissingParameterError {
		parameterName: parameterName,
	}
}

func (e *MissingParameterError) Error() string {
	return fmt.Sprintf("could not find parameter \"%s\"", e.parameterName)
}

func getRequest(r *http.Request) (v interface{}, err error) {
	var bodyBytes []byte
	_, err = r.Body.Read(bodyBytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func writeResponse(w http.ResponseWriter, res interface{}, status int) {

	resBytes, err := json.Marshal(res)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(resBytes)
	if err != nil {
		writeError(w, err)
		return
	}
}
func writeEmptyResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func writeError(w http.ResponseWriter, err error) {

	// Todo: Handle different kinds of errors
	switch err.(type) {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}