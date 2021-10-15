package modules

import (
	"fmt"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/config"
	"github.com/yukitsune/chameleon/internal/api/db"
	mongo2 "github.com/yukitsune/chameleon/internal/api/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
)

type DatabaseModule struct {
	Config *config.DbConfig
}

func (m *DatabaseModule) Register(cb camogo.ContainerBuilder) error {
	var err error

	// Config
	err = cb.RegisterInstance(m.Config)
	if err != nil {
		return err
	}

	// Database
	err = cb.RegisterFactory(func(cfg *config.DbConfig) (db.ConnectionWrapper, error) {

		uri := fmt.Sprintf(
			"mongodb://%s:%d/%s",
			url.QueryEscape(cfg.Host),
			cfg.Port,
			url.QueryEscape(cfg.Database))

		creds := options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		}
		opts := options.Client().ApplyURI(uri).SetAuth(creds)

		client, err := mongo.NewClient(opts)
		if err != nil {
			return nil, err
		}

		wrapper := mongo2.NewMongoConnectionWrapper(client, cfg.Database)
		return wrapper, nil
	},
		camogo.TransientLifetime)
	if err != nil {
		return err
	}

	return nil
}
