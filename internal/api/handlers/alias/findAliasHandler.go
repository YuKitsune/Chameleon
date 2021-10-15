package alias

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/db"
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/repository"
	"regexp"
)

type FindAliasHandler struct {
	ctx context.Context
	db db.ConnectionWrapper
	log *logrus.Logger
}

func NewFindAliasHandler(ctx context.Context, db db.ConnectionWrapper, log *logrus.Logger) *FindAliasHandler {
	return &FindAliasHandler{ctx, db, log}
}

func (handler *FindAliasHandler) Handle(req *model.FindAliasRequest) (*model.Alias, error) {

	var allAliasesForRecipient []model.Alias
	err := handler.db.InConnection(handler.ctx, func (ctx context.Context, ds repository.DataSource) error {
		collection := ds.Collection("alias")
		return collection.FindAll(handler.ctx, repository.Filter{
			"username": req.Recipient,
		}, &allAliasesForRecipient)
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
