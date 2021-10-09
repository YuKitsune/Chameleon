package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/config"
	"github.com/yukitsune/chameleon/internal/api/middleware"
	"github.com/yukitsune/chameleon/internal/api/modules"
	"github.com/yukitsune/chameleon/internal/api/routers"
	"github.com/yukitsune/chameleon/internal/log"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config    *config.Config
	svr       *http.Server
	container camogo.Container
	log       log.ChameleonLogger
}

func NewChameleonApiServer(config *config.Config, logger log.ChameleonLogger) (*ChameleonApiServer, error) {
	api := &ChameleonApiServer{
		config: config,
		log:    logger,
	}

	c, err := buildContainer(config.Database, logger)
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
	api.log.Infof("Listening on TCP %s", api.svr.Addr)
	api.log.Warningln("TLS not enabled")
	return api.svr.ListenAndServe()
}

func (api *ChameleonApiServer) StartTLS() error {
	api.log.Infof("Listening on TCP %s with TLS", api.svr.Addr)
	return api.svr.ListenAndServeTLS(api.config.CertFile, api.config.KeyFile)
}

func buildContainer(dbConfig *config.DbConfig, logger log.ChameleonLogger) (camogo.Container, error) {
	cb := camogo.NewBuilder()

	var err error

	// Context
	// Todo: Gonna need a different way of providing context
	err = cb.RegisterFactory(func () context.Context {
		return context.TODO()
	},
		camogo.TransientLifetime)
	if err != nil {
		return nil, err
	}

	// Logger
	err = cb.RegisterFactory(func () log.ChameleonLogger {
		return logger
	},
		camogo.TransientLifetime)
	if err != nil {
		return nil, err
	}

	// Database
	err = cb.RegisterModule(&modules.DatabaseModule{Config: dbConfig})
	if err != nil {
		return nil, err
	}

	// Mediator
	err = cb.RegisterModule(&modules.MediatorHandlerModule{})
	if err != nil {
		return nil, err
	}

	c := cb.Build()
	return c, nil
}

func makeHandler(container camogo.Container) http.Handler {
	m := mux.NewRouter()

	// Middleware
	// Panic recovery
	panicRecoveryMiddleware := middleware.NewPanicRecovery(container)
	m.Use(panicRecoveryMiddleware.Middleware)

	// Endpoints
	routers.NewAliasRouter(m.PathPrefix("/alias").Subrouter(), container)

	return m
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}
