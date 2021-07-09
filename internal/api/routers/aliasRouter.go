package routers

import (
	"github.com/gorilla/mux"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"net/http"
	"reflect"
)

type AliasRouter struct {
	r *mux.Router
	container camogo.Container
	hasBeenSetUp bool
}

func NewAliasRouter(r *mux.Router, c camogo.Container) *AliasRouter {
	router :=  &AliasRouter{
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

	router.r.HandleFunc("/create", router.Create)
	router.r.HandleFunc("", router.Read).
		Queries(
			"sender", "{sender}",
			"recipient", "{recipient}")
	router.r.HandleFunc("/update", router.Update)
	router.r.HandleFunc("/delete", router.Delete)

	router.hasBeenSetUp = true
}

func (router *AliasRouter) Create(w http.ResponseWriter, r *http.Request) {

	// Parse the request JSON
	req, err := getRequest(r)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Cast the request so we can send it over the mediator
	createRequest, ok := req.(model.CreateAliasRequest)
	if !ok {
		responseWriterHelpers.WriteError(w, model.CastFailedErr(reflect.TypeOf(&model.CreateAliasRequest{})))
		return
	}

	// Send the request through the mediator
	var res interface{}
	err = router.container.Resolve(func (mediator *mediator.Mediator) error {
		res, err = mediator.Send(createRequest)
		return err
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
	var res interface{}
	var err error
	err = router.container.Resolve(func (mediator *mediator.Mediator) error {
		res, err = mediator.Send(req)
		return err
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
	req, err := getRequest(r)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Cast the request so we can send it over the mediator
	updateRequest, ok := req.(model.UpdateAliasRequest)
	if !ok {
		responseWriterHelpers.WriteError(w, model.CastFailedErr(reflect.TypeOf(&model.UpdateAliasRequest{})))
		return
	}

	// Send the request through the mediator
	var res interface{}
	err = router.container.Resolve(func (mediator *mediator.Mediator) error {
		res, err = mediator.Send(updateRequest)
		return err
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
	req, err := getRequest(r)
	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Cast the request so we can send it over the mediator
	deleteRequest, ok := req.(model.DeleteAliasRequest)
	if !ok {
		responseWriterHelpers.WriteError(w, model.CastFailedErr(reflect.TypeOf(&model.DeleteAliasRequest{})))
		return
	}

	// Send the request through the mediator
	var res interface{}
	err = router.container.Resolve(func (mediator *mediator.Mediator) error {
		res, err = mediator.Send(deleteRequest)
		return err
	})

	if err != nil {
		responseWriterHelpers.WriteError(w, err)
		return
	}

	// Write the response
	responseWriterHelpers.WriteEmptyResponse(w, http.StatusOK)
}
