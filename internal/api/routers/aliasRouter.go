package routers

import (
	"github.com/gorilla/mux"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"net/http"
)

type AliasRouter struct {
	r            *mux.Router
	container    camogo.Container
	hasBeenSetUp bool
}

func NewAliasRouter(r *mux.Router, c camogo.Container) *AliasRouter {
	router := &AliasRouter{
		r:         r,
		container: c,
	}

	router.setup()
	return router
}

func (router *AliasRouter) setup() {
	if router.hasBeenSetUp {
		return
	}

	router.r.HandleFunc("", router.Create).Methods("POST")
	router.r.HandleFunc("", router.Read).Methods("GET").
		Queries(
			"sender", "{sender}",
			"recipient", "{recipient}")
	router.r.HandleFunc("", router.Update).Methods("PUT")
	router.r.HandleFunc("", router.Delete).Methods("DELETE")

	router.hasBeenSetUp = true
}

func (router *AliasRouter) Create(w http.ResponseWriter, r *http.Request) {

	// Parse the request JSON
	var createRequest model.CreateAliasRequest
	err := getRequest(r, &createRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	res, err := router.container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&createRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusCreated)
}

func (router *AliasRouter) Read(w http.ResponseWriter, r *http.Request) {

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
	req := &model.GetAliasRequest{
		Sender:    sender,
		Recipient: recipient,
	}

	// Send the request through the mediator
	res, err := router.container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(req)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusOK)
}

func (router *AliasRouter) Update(w http.ResponseWriter, r *http.Request) {

	// Parse the request JSON
	var updateRequest model.UpdateAliasRequest
	err := getRequest(r, &updateRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	res, err := router.container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&updateRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteResponse(w, res, http.StatusOK)
}

func (router *AliasRouter) Delete(w http.ResponseWriter, r *http.Request) {

	// Parse the request JSON
	var deleteRequest model.DeleteAliasRequest
	err := getRequest(r, &deleteRequest)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Send the request through the mediator
	_, err = router.container.ResolveWithResult(func(mediator mediator.Mediator) (interface{}, error) {
		return mediator.Send(&deleteRequest)
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteEmptyResponse(w, http.StatusOK)
}
