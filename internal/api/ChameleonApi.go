package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yukitsune/chameleon/internal/api/handlers"
	"github.com/yukitsune/chameleon/internal/api/mediatorHandlers"
	"github.com/yukitsune/chameleon/internal/api/routers"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config    *ApiConfig
	svr       *http.Server
	container ioc.Container
}

func NewChameleonApiServer(config *ApiConfig, logger log.ChameleonLogger) (*ChameleonApiServer, error) {
	api := &ChameleonApiServer{
		config: config,
	}

	c, err := makeContainer(config.Database, logger)
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

func makeContainer(dbConfig *DbConfig, logger log.ChameleonLogger) (ioc.Container, error) {
	c := ioc.NewGolobbyContainer()
	var err error

	err = c.RegisterSingletonInstance(logger)
	if err != nil {
		return nil, err
	}

	err = c.RegisterSingletonInstance(dbConfig)
	if err != nil {
		return nil, err
	}

	err = c.RegisterTransientFactory(func (cfg *DbConfig) *gorm.DB {
		// Todo: Handle error
		db, _ := gorm.Open(postgres.Open(cfg.ConnectionString()), &gorm.Config{})
		return db
	})
	if err != nil {
		return nil, err
	}

	err = c.RegisterTransientFactory(mediatorHandlers.NewValidateHandler)
	if err != nil {
		return nil, err
	}

	err = c.RegisterTransientFactory(mediatorHandlers.NewMailHandler)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func makeHandler(container ioc.Container) http.Handler {
	m := mux.NewRouter()

	routers.NewAliasRouter(m.Path("/alias").Subrouter())

	return m
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}
