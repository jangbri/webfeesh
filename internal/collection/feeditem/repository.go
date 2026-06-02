package feeditem

import (
	"context"
)

type Repository interface {
	GetLatest(ctx context.Context, cID int64) ([]*CollectionFeedItem, error)
}
