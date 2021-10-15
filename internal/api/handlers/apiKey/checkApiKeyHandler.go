package apiKey

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CheckApiKeyHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log *logrus.Logger
}

func NewCheckApiKeyHandler(ctx context.Context, db *db.MongoConnectionWrapper, log *logrus.Logger) *CheckApiKeyHandler {
	return &CheckApiKeyHandler{ctx, db, log}
}

func (h *CheckApiKeyHandler) Handle(req *model.CheckApiKeyRequest) (bool, error) {

	var foundApiKey *model.ApiKey
	err := h.db.InConnection(h.ctx, func (ctx context.Context, db *mongo.Database) error {
		collection := db.Collection(model.ApiKeyCollectionName)
		res := collection.FindOne(h.ctx, bson.M{
			"value": req.Value,
		})

		err := res.Err()
		if err != nil {
			return res.Err()
		}

		return res.Decode(&foundApiKey)
	})

	if err != nil {
		return false, err
	}

	if foundApiKey != nil && foundApiKey.Scopes.ContainsAll(req.Scopes) {
		return true, nil
	}

	return false, nil
}
