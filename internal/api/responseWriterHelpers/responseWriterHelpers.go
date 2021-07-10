package responseWriterHelpers

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, res interface{}, status int) {

	resBytes, err := json.Marshal(res)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(resBytes)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func WriteEmptyResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func WriteError(w http.ResponseWriter, err error) {

	// Todo: Handle different kinds of errors
	switch err.(type) {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
