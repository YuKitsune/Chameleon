package apiKey

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	mongo2 "github.com/yukitsune/chameleon/internal/api/db/mongo"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/repository"
)

type CheckApiKeyHandler struct {
	ctx context.Context
	db db.ConnectionWrapper
	log *logrus.Logger
}

func NewCheckApiKeyHandler(ctx context.Context, db *mongo2.MongoConnectionWrapper, log *logrus.Logger) *CheckApiKeyHandler {
	return &CheckApiKeyHandler{ctx, db, log}
}

func (h *CheckApiKeyHandler) Handle(req *model.CheckApiKeyRequest) (bool, error) {

	var foundApiKey *model.ApiKey
	err := h.db.InConnection(h.ctx, func (ctx context.Context, ds repository.DataSource) error {
		collection := ds.Collection(model.ApiKeyCollectionName)
		return collection.FindOne(h.ctx, repository.Filter{
			"value": req.Value,
		}, &foundApiKey)
	})

	if err != nil {
		return false, err
	}

	if foundApiKey != nil && foundApiKey.Scopes.ContainsAll(req.Scopes) {
		return true, nil
	}

	return false, nil
}
