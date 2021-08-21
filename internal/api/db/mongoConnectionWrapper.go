package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnectionWrapper struct {
	Client   *mongo.Client
	Database string
}

func (w *MongoConnectionWrapper) InConnection(ctx context.Context, fn func (ctx context.Context, db *mongo.Database) error) (err error) {
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
	err = fn(ctx, db)
	return err
}
