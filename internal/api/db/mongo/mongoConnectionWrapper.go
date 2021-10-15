package mongo

import (
	"context"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/repository"
	mongo2 "github.com/yukitsune/chameleon/internal/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnectionWrapper struct {
	Client   *mongo.Client
	Database string
}

func NewMongoConnectionWrapper(c *mongo.Client, db string) db.ConnectionWrapper {
	return &MongoConnectionWrapper{c, db}
}

func (w *MongoConnectionWrapper) InConnection(ctx context.Context, fn func (ctx context.Context, ds repository.DataSource) error) (err error) {
	err = w.Client.Connect(ctx)
	if err != nil {
		return nil
	}

	defer func() {
		err = w.Client.Disconnect(ctx)
	}()

	err = w.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db := w.Client.Database(w.Database)
	ds := mongo2.NewMongoDataSource(db)
	err = fn(ctx, ds)
	return err
}
