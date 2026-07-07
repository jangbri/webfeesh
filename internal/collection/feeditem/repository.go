package feeditem

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, c CollectionFeedItem) error
	GetLatest(ctx context.Context, cID int64) ([]CollectionFeedItem, error)
}
