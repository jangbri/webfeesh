package feed

import (
	"context"
	"log/slog"

	"github.com/jangbri/webfeesh/internal/item"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, feed *Feed) (*Feed, error) {
	err := validateFeedLink(feed.Link)
	if err != nil {
		slog.Error("invalid feed link")
		return nil, err
	}
	return s.repo.Create(ctx, feed)
}

func (s *Service) Update(ctx context.Context, feed *Feed) (*Feed, error) {
	err := validateFeedLink(feed.Link)
	if err != nil {
		slog.Error("invalid feed link")
		return nil, err
	}
	return s.repo.Update(ctx, feed)
}

func (s *Service) Delete(ctx context.Context, feed *Feed) error {
	return s.repo.Delete(ctx, feed)
}

func (s *Service) GetFeedItems(ctx context.Context, feed *Feed) ([]*item.Item, error) {
	return s.repo.GetFeedItems(ctx, feed)
}

func (s *Service) GetAllFeeds(ctx context.Context) ([]*Feed, error) {
	return s.repo.GetAllFeeds(ctx)
}
