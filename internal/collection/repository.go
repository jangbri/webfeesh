package collection

import (
	"context"

	"github.com/jangbri/webfeesh/internal/feed"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Collection, error)
	Create(ctx context.Context, collection Collection) error
	Update(ctx context.Context, collection Collection) error
	Delete(ctx context.Context, collection Collection) error

	GetCollectionFeeds(ctx context.Context, collection Collection) ([]feed.Feed, error)
}
