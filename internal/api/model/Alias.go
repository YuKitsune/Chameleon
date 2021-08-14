package model

import (
	"regexp"
)

type Alias struct {
	Id string `bson:"_id"`
	Username               string
	SenderWhitelistPattern string
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

type GetAliasRequest struct {
	Sender    string
	Recipient string
}

type UpdateAliasRequest struct {
	Alias
}

type DeleteAliasRequest struct {
	Alias
}

// Todo: Temporary until I can be bothered to allow camogo to return primitives

type DeleteAliasResponse struct {
	Deleted bool
}