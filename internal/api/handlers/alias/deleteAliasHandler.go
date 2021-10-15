package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteAliasHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log *logrus.Logger
}

func NewDeleteAliasHandler(ctx context.Context, db *db.MongoConnectionWrapper, log *logrus.Logger) *DeleteAliasHandler {
	return &DeleteAliasHandler{ctx, db, log}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (bool, error) {
	deleted := false
	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, db *mongo.Database) error {
		collection := db.Collection("alias")
		_, err := collection.DeleteOne(handler.ctx, bson.M{"_id": req.Alias.Id})
		if err != nil {
			return err
		}

		deleted = true
		return nil
	})
	if err != nil {
		return false, err
	}

	return deleted, nil
}
