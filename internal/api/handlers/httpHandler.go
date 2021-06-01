package handlers

import (
	"encoding/json"
	"net/http"
)

type HttpHandlerWrapper struct {
	inner Handler
}

func NewHttpHandler(inner Handler) HttpHandlerWrapper {
	return HttpHandlerWrapper{
		inner: inner,
	}
}

func (h HttpHandlerWrapper) Handle(v interface{}) (interface{}, error) {
	return h.inner.Handle(v)
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
	res, err := h.Handle(req)

	// Todo: Any errors in this block should be added to the one from the above Handle call
	if res != nil {
		resBytes, err := json.Marshal(res)
		if err != nil {
			_, err = w.Write(resBytes)
		}
	}

	// Todo: Handle different error types
	switch err.(type) {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
