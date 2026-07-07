package feed

import (
	"context"

	"github.com/jangbri/webfeesh/internal/item"
)

type Repository interface {
	Create(ctx context.Context, feed Feed) (Feed, error)
	Update(ctx context.Context, feed Feed) (Feed, error)
	Delete(ctx context.Context, feed Feed) error

	GetFeedItems(ctx context.Context, feed Feed) ([]item.Item, error)

	GetAllFeeds(ctx context.Context) ([]Feed, error)
}
