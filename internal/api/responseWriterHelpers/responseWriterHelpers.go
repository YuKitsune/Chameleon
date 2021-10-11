package responseWriterHelpers

import (
	"encoding/json"
	"fmt"
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/pkg/mediator"
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

func WriteUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func WriteBadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(message))
}

func WriteBadRequestf(w http.ResponseWriter, format string, s ...interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(format, s...)))
}

func WriteError(w http.ResponseWriter, err error) {

	// Todo: Handle different kinds of errors
	switch err.(type) {

	// Unwrap mediator errors
	case *mediator.Error:
		mediatorError := err.(*mediator.Error)
		WriteError(w, mediatorError.Unwrap())
		break

	case *errors.EntityNotFoundError:
		w.WriteHeader(http.StatusNotFound)
		break

	default:
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}
