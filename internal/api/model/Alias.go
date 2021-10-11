package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

type Alias struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username               string `bson:"username" json:"username"`
	SenderWhitelistPattern string `bson:"senderWhitelistPattern" json:"senderWhitelistPattern"`
}

func (a *Alias) SenderIsAllowed(sender string) (*bool, error) {
	r, err := regexp.Compile(a.SenderWhitelistPattern)
	if err != nil {
		return nil, err
	}

	match := r.MatchString(sender)
	return &match, nil
}

type CreateAliasRequest struct {
	Alias
}

type FindAliasRequest struct {
	Sender    string
	Recipient string
}

type UpdateAliasRequest struct {
	Alias
}

type DeleteAliasRequest struct {
	Alias
}
