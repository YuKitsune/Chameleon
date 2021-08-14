package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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

	db := w.Client.Database(w.Database)
	err = fn(ctx, db)
	return err
}
