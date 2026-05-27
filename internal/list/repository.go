package list

import (
	"context"

	"github.com/jangbri/webfeesh-be/internal/feed"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*List, error)
	Create(ctx context.Context, list *List) error
	Update(ctx context.Context, list *List) error
	Delete(ctx context.Context, list *List) error

	GetListFeeds(ctx context.Context, list *List) ([]*feed.Feed, error)
}
