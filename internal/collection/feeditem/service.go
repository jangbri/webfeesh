package feeditem

import (
	"context"
	"log/slog"

	"github.com/jangbri/webfeesh/internal/collection"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetLatest(
	ctx context.Context,
	collection *collection.Collection,
) ([]*CollectionFeedItem, error) {
	items, err := s.repo.GetLatest(ctx, collection)
	if err != nil {
		slog.Error("error when retrieving collection's aggregated feed items")
		return nil, err
	}

	return items, nil
}
