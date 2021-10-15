package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/repository"
)

type UpdateAliasHandler struct {
	ctx context.Context
	db db.ConnectionWrapper
	log *logrus.Logger
}

func NewUpdateAliasHandler(ctx context.Context, db db.ConnectionWrapper, log *logrus.Logger) *UpdateAliasHandler {
	return &UpdateAliasHandler{ctx, db, log}
}

func (handler *UpdateAliasHandler) Handle(req *model.UpdateAliasRequest) (*model.Alias, error) {

	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, ds repository.DataSource) error {
		collection := ds.Collection("alias")
		return collection.UpdateById(handler.ctx, req.Alias.Id, req.Alias)
	})
	if err != nil {
		return nil, err
	}

	return &req.Alias, nil
}
