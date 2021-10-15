package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteAliasHandler struct {
	ctx context.Context
	db db.ConnectionWrapper
	log *logrus.Logger
}

func NewDeleteAliasHandler(ctx context.Context, db db.ConnectionWrapper, log *logrus.Logger) *DeleteAliasHandler {
	return &DeleteAliasHandler{ctx, db, log}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (bool, error) {
	deleted := false
	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, ds repository.DataSource) error {
		collection := ds.Collection("alias")
		err := collection.DeleteById(handler.ctx, bson.M{"_id": req.Alias.Id})
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
