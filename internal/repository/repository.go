package repository

import "context"

type Filter map[string]interface{}

type DataSource interface {
	Collection(string) Repository
}

type Repository interface {
	Add(ctx context.Context, doc interface{}) error

	Count(ctx context.Context, filter Filter) (int64, error)
	FindAll(ctx context.Context, filter Filter, result interface{}) error
	FindOne(ctx context.Context, filter Filter, result interface{}) error

	UpdateById(ctx context.Context, id interface{}, doc interface{}) error
	DeleteById(ctx context.Context, id interface{}) error
}