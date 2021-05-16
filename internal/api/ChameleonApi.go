package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yukitsune/chameleon/internal/api/handlers"
	"github.com/yukitsune/chameleon/internal/log"
	"go.uber.org/dig"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config    *ApiConfig
	svr       *http.Server
	container *dig.Container
}

func NewChameleonApiServer(config *ApiConfig, logger log.ChameleonLogger) (*ChameleonApiServer, error) {
	api := &ChameleonApiServer{
		config: config,
	}

	c, err := makeContainer(logger)
	if err != nil {
		return nil, err
	}

	h := makeHandler(c)
	api.svr = &http.Server{
		Addr: api.config.GetAddress(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}

	return api, nil
}

func (api *ChameleonApiServer) Start() error {
	return api.svr.ListenAndServe()
}

func (api *ChameleonApiServer) StartTLS() error {
	return api.svr.ListenAndServeTLS(api.config.CertFile, api.config.KeyFile)
}

func makeContainer(logger log.ChameleonLogger) (*dig.Container, error) {
	c := dig.New()
	var err error

	err = c.Provide(func() log.ChameleonLogger { return logger })
	if err != nil {
		return nil, err
	}

	err = c.Provide(handlers.NewValidateHandler)
	if err != nil {
		return nil, err
	}

	err = c.Provide(handlers.NewMailHandler)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func makeHandler(container *dig.Container) http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/validate", func(writer http.ResponseWriter, request *http.Request) {
		_ = container.Invoke(func(handler *handlers.ValidateHandler) {
			handler.Handle(writer, request)
		})
	})

	m.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		_ = container.Invoke(func(handler *handlers.MailHandler) {
			handler.Handle(writer, request)
		})
	})

	return m
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}
