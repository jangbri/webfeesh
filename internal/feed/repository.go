package feed

import (
	"context"

	"github.com/jangbri/webfeesh-be/internal/item"
)

type Repository interface {
	Create(ctx context.Context, feed *Feed) (*Feed, error)
	Update(ctx context.Context, feed *Feed) (*Feed, error)
	Delete(ctx context.Context, feed *Feed) error

	GetFeedItems(ctx context.Context, feed *Feed) ([]*item.Item, error)
}
