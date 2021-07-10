package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/handlers/alias"
	"github.com/yukitsune/chameleon/internal/api/middleware"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/api/routers"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ChameleonApiServer struct {
	config    *Config
	svr       *http.Server
	container camogo.Container
	log log.ChameleonLogger
}

func NewChameleonApiServer(config *Config, logger log.ChameleonLogger) (*ChameleonApiServer, error) {
	api := &ChameleonApiServer{
		config: config,
		log: logger,
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
	api.log.Infof("Listening on TCP %s", api.svr.Addr)
	api.log.Warningln("TLS not enabled")
	return api.svr.ListenAndServe()
}

func (api *ChameleonApiServer) StartTLS() error {
	api.log.Infof("Listening on TCP %s with TLS", api.svr.Addr)
	return api.svr.ListenAndServeTLS(api.config.CertFile, api.config.KeyFile)
}

func makeContainer(dbConfig *DbConfig, logger log.ChameleonLogger) (camogo.Container, error) {
	c := camogo.New()

	err := c.Register(func (r *camogo.Registrar) error {
		var err error

		// Logger
		err = r.RegisterInstance(logger)
		if err != nil {
			return err
		}

		// Database config
		err = r.RegisterInstance(dbConfig)
		if err != nil {
			return err
		}

		// Database
		err = r.RegisterFactory(func (cfg *DbConfig) (*gorm.DB, error) {
			db, err := gorm.Open(postgres.Open(cfg.ConnectionString()), &gorm.Config{})
			if err != nil {
				return nil, err
			}

			err = db.AutoMigrate(&model.User{}, &model.Alias{})
			if err != nil {
				return nil, err
			}

			return db, nil
		},
		camogo.TransientLifetime)
		if err != nil {
			return err
		}

		// The container itself
		err = r.RegisterFactory(func () camogo.Container {
			return c
		},
		camogo.TransientLifetime)
		if err != nil {
			return err
		}

		// Mediator
		err = r.RegisterFactory(makeMediator, camogo.SingletonLifetime)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

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

func makeMediator(container camogo.Container) (*mediator.Mediator, error) {
	m := mediator.New(container)
	var err error

	err = m.AddRequestHandlerFactory(alias.NewCreateAliasHandler)
	if err != nil {
		return nil, err
	}

	err = m.AddRequestHandlerFactory(alias.NewReadAliasHandler)
	if err != nil {
		return nil, err
	}

	err = m.AddRequestHandlerFactory(alias.NewUpdateAliasHandler)
	if err != nil {
		return nil, err
	}

	err = m.AddRequestHandlerFactory(alias.NewDeleteAliasHandler)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (api *ChameleonApiServer) Shutdown() error {
	return api.svr.Shutdown(context.Background())
}
