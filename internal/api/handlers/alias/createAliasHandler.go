package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

type CreateAliasHandler struct {
	ctx context.Context
	db *db.MongoConnectionWrapper
	log *logrus.Logger
}

func NewCreateAliasHandler(ctx context.Context, db *db.MongoConnectionWrapper , log *logrus.Logger) *CreateAliasHandler {
	return &CreateAliasHandler{ctx, db, log}
}

func (handler *CreateAliasHandler) Handle(req *model.CreateAliasRequest) (*model.Alias, error) {
	alias := req.Alias

	// Ensure the request is valid

	// Todo: User must exist and be the user submitting the request
	//var userCount int64
	//handler.db.Where(&model.User{}, alias.UserID).Count(&userCount)
	//if userCount == 0 {
	//	return nil, errors.NewEntityInvalidError(&alias.Model, "user does not exist")
	//}

	// Recipient must not be empty
	if len(alias.Username) == 0 {
		return nil, errors.NewEntityInvalidError(&alias, "username must not be empty")
	}

	// Sender whitelist pattern must be valid regex
	_, err := regexp.Compile(alias.SenderWhitelistPattern)
	if err != nil {
		return nil, errors.NewEntityInvalidErrorFromErr(&alias, err)
	}

	// Ensure no duplicate entries exist
	err = handler.db.InConnection(handler.ctx, func (ctx context.Context, db *mongo.Database) error {
		collection := db.Collection("alias")
		dupes, err := collection.CountDocuments(ctx, bson.M{
			"Username": alias.Username,
			"SenderWhitelistPattern": alias.SenderWhitelistPattern,
		})
		if err != nil {
			return err
		}
		if dupes != 0 {
			return errors.NewEntityExistsError(&alias)
		}

		// All good, create the record
		_, err = collection.InsertOne(handler.ctx, alias)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &alias, nil
}
