package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const ApiKeyCollectionName = "ApiKeys"

type ApiKey struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OwnerId primitive.ObjectID `bson:"ownerId" json:"ownerId"`
	Value string `bson:"value" json:"value"`
	Scopes Scopes `bson:"scopes" json:"scopes"`
}

type CreateApiKeyRequest struct {
	Scopes Scopes
}

type CheckApiKeyRequest struct {
	Value string
	Scopes Scopes
}