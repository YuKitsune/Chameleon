package alias

import (
	"context"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteAliasHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log log.ChameleonLogger
}

func NewDeleteAliasHandler(ctx context.Context, db *db.MongoConnectionWrapper, log log.ChameleonLogger) *DeleteAliasHandler {
	return &DeleteAliasHandler{ctx, db, log}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (*model.DeleteAliasResponse, error) {
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
		return nil, err
	}

	return &model.DeleteAliasResponse{Deleted: deleted}, nil
}
