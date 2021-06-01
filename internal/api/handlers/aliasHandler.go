package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type AliasHandler struct {
	router *mux.Router
}

func (h *AliasHandler) add(w http.ResponseWriter, r *http.Request) {

}

func (h *AliasHandler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *AliasHandler) modify(w http.ResponseWriter, r *http.Request) {

}

func (h *AliasHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *AliasHandler) setupRouting() {
	h.router.HandleFunc("/", h.get).Methods("GET")
	h.router.HandleFunc("/add", h.add).Methods("POST")
	h.router.HandleFunc("/modify", h.modify).Methods("PUT")
	h.router.HandleFunc("/delete", h.delete).Methods("DELETE")
}