package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/config"
	"github.com/yukitsune/chameleon/internal/api/middleware"
	"github.com/yukitsune/chameleon/internal/api/modules"
	"github.com/yukitsune/chameleon/internal/api/routers"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config    *config.Config
	svr       *http.Server
	container camogo.Container
	log       *logrus.Logger
}

func NewChameleonApiServer(config *config.Config, logger *logrus.Logger) (*ChameleonApiServer, error) {
	api := &ChameleonApiServer{
		config: config,
		log:    logger,
	}

	c, err := buildContainer(config.Database, logger)
	if err != nil {
		return nil, err
	}

	h := buildHandler(c)
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

func buildContainer(dbConfig *config.DbConfig, logger *logrus.Logger) (camogo.Container, error) {
	cb := camogo.NewBuilder()

	var err error

	// Logger
	err = cb.RegisterFactory(func () *logrus.Logger {
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
	//err = cb.RegisterModule(&modules.MediatorHandlerModule{})
	//if err != nil {
	//	return nil, err
	//}

	c := cb.Build()
	return c, nil
}

func buildHandler(container camogo.Container) http.Handler {
	r := mux.NewRouter()

	configureMiddleware(r, container)
	configureEndpoints(r)

	return r
}

func configureMiddleware(r *mux.Router, container camogo.Container) {

	// Container injection
	containerInjectionMiddleware := middleware.NewContainerInjectionMiddleware(container)
	r.Use(containerInjectionMiddleware.Middleware)

	r.Use(middleware.PanicRecovery)
}

func configureEndpoints(r *mux.Router) {
	mountRouter(r, "/alias", routers.AliasRouter)
}

func mountRouter(r *mux.Router, path string, handler func (*mux.Router)) {
	handler(r.PathPrefix(path).Subrouter())
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}
