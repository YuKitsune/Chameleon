package db

import (
	"context"
	"github.com/yukitsune/chameleon/internal/repository"
)

// Todo: Merge with repository package
type ConnectionWrapper interface {
	InConnection(ctx context.Context, fn func (ctx context.Context, ds repository.DataSource) error) (err error)
}
