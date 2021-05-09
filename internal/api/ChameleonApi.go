package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yukitsune/chameleon/internal/api/handlers"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config *ApiConfig
	svr *http.Server
	router *mux.Router

	validateHandler handlers.ValidateHandler
	mailHandler handlers.MailHandler
}

func NewChameleonApiServer(config *ApiConfig) *ChameleonApiServer {
	api := &ChameleonApiServer{
		config: config,
	}

	h := api.makeHandler()
	api.svr = &http.Server{
		Addr:         api.config.GetAddress(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: h,
	}

	return api
}

func (api *ChameleonApiServer) Start() error {
	return api.svr.ListenAndServe()
}

func (api *ChameleonApiServer) StartTLS() error {
	return api.svr.ListenAndServeTLS(api.config.CertFile, api.config.KeyFile)
}

func (api *ChameleonApiServer) makeHandler() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc("/validate", api.validateHandler.Handle)
	m.HandleFunc("/handle", api.mailHandler.Handle)
	return m
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}