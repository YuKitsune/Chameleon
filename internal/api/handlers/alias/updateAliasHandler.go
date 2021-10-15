package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAliasHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log *logrus.Logger
}

func NewUpdateAliasHandler(ctx context.Context, db *db.MongoConnectionWrapper, log *logrus.Logger) *UpdateAliasHandler {
	return &UpdateAliasHandler{ctx, db, log}
}

func (handler *UpdateAliasHandler) Handle(req *model.UpdateAliasRequest) (*model.Alias, error) {

	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, db *mongo.Database) error {
		collection := db.Collection("alias")
		_, err := collection.UpdateByID(handler.ctx, req.Alias.Id, req.Alias)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &req.Alias, nil
}
