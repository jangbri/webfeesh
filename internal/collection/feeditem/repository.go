package feeditem

import (
	"context"

	"github.com/jangbri/webfeesh/internal/collection"
)

type Repository interface {
	GetLatest(ctx context.Context, collection *collection.Collection) ([]*CollectionFeedItem, error)
}
