package routers

import (
	"github.com/gorilla/mux"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"net/http"
)

type AliasRouter struct {
	r *mux.Router
	mediator *mediator.Mediator
}

func NewAliasRouter(r *mux.Router) {

}

func (router *AliasRouter) Create(w http.ResponseWriter, r *http.Request) {
	res, err := router.mediator.Send()
}

