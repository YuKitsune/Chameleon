package alias

import (
	"context"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

type ReadAliasHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log log.ChameleonLogger
}

func NewReadAliasHandler(ctx context.Context, db *db.MongoConnectionWrapper, log log.ChameleonLogger) *ReadAliasHandler {
	return &ReadAliasHandler{ctx, db, log}
}

func (handler *ReadAliasHandler) Handle(req *model.GetAliasRequest) (*model.Alias, error) {

	var allAliasesForRecipient []model.Alias
	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, db *mongo.Database) error {
		collection := db.Collection("alias")
		cur, err := collection.Find(handler.ctx, bson.M{
			"username": req.Recipient,
		})
		if err != nil {
			return err
		}

		err = cur.All(handler.ctx, &allAliasesForRecipient)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, alias := range allAliasesForRecipient {
		isMatch, err := regexp.MatchString(alias.SenderWhitelistPattern, req.Sender)
		if err != nil {
			return nil, err
		}

		if isMatch {
			return &alias, nil
		}
	}

	return nil, errors.NewEntityNotFoundError(req)
}
