package responseWriterHelpers

import (
	"encoding/json"
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
