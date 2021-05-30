package handlers

import (
	"encoding/json"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"net/http"
)

type handlerFunc func (ioc.Container, interface{}) error

type HttpHandlerWrapper struct {
	container ioc.Container
	handlerFn handlerFunc
}

func NewHttpHandler(container ioc.Container, handlerFn handlerFunc) HttpHandlerWrapper {
	return HttpHandlerWrapper{
		container: container,
		handlerFn: handlerFn,
	}
}

func (h HttpHandlerWrapper) Handle(v interface{}) error {
	return h.handlerFn(h.container, v)
}

func (h HttpHandlerWrapper) HandleHttp(w http.ResponseWriter, r *http.Request) {

	// Read the body as JSON
	var bodyBytes []byte
	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var req interface{}
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Handle the request
	err = h.Handle(req)

	// Todo: Handle different error types
	switch err.(type) {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
