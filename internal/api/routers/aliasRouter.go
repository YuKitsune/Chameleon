package routers

import (
	"github.com/gorilla/mux"
	"github.com/yukitsune/chameleon/internal/api/context"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"net/http"
)

func AliasRouter(r *mux.Router) {
	r.HandleFunc("", Create).Methods("POST")
	r.HandleFunc("", Find).Methods("GET").
		Queries(
			"sender", "{sender}",
			"recipient", "{recipient}")
	r.HandleFunc("", Update).Methods("PUT")
	r.HandleFunc("", Delete).Methods("DELETE")
}

func Create(w http.ResponseWriter, r *http.Request) {

	container, err := context.Container(r.Context())
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Parse the request JSON
	var createRequest model.CreateAliasRequest
	err = getRequest(r, &createRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	res, err := container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&createRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusCreated)
}

func Find(w http.ResponseWriter, r *http.Request) {

	container, err := context.Container(r.Context())
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Get the parameters from the URL
	vars := mux.Vars(r)
	sender, ok := vars["sender"]
	if !ok {
		responseWriterHelpers.WriteError(w, MissingParameterErr("sender"))
		return
	}

	recipient, ok := vars["recipient"]
	if !ok {
		responseWriterHelpers.WriteError(w, MissingParameterErr("recipient"))
		return
	}

	// Build the request
	req := &model.FindAliasRequest{
		Sender:    sender,
		Recipient: recipient,
	}

	// Send the request through the mediator
	res, err := container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(req)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {

	container, err := context.Container(r.Context())
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Parse the request JSON
	var updateRequest model.UpdateAliasRequest
	err = getRequest(r, &updateRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	res, err := container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&updateRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	container, err := context.Container(r.Context())
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Parse the request JSON
	var deleteRequest model.DeleteAliasRequest
	err = getRequest(r, &deleteRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	_, err = container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&deleteRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteEmptyResponse(w, http.StatusOK)
}
