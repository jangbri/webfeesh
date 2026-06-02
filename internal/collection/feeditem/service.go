package feeditem

import (
	"context"
	"log/slog"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetLatest(ctx context.Context, cID int64) ([]*CollectionFeedItem, error) {
	items, err := s.repo.GetLatest(ctx, cID)
	if err != nil {
		slog.Error("error when retrieving collection's aggregated feed items")
		return nil, err
	}

	return items, nil
}
