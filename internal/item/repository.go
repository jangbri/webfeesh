package item

import "context"

type Repository interface {
	UpsertMany(ctx context.Context, items []*Item) error
}
