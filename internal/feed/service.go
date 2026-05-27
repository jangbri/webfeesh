package feed

import (
	"context"

	"github.com/jangbri/webfeesh-be/internal/item"
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
	// TODO: validate the feed link to make sure its a valid feed link
	return s.repo.Create(ctx, feed)
}

func (s *Service) Update(ctx context.Context, feed *Feed) (*Feed, error) {
	// TODO: validate the feed link to make sure its a valid feed link
	return s.repo.Update(ctx, feed)
}

func (s *Service) Delete(ctx context.Context, feed *Feed) error {
	return s.repo.Delete(ctx, feed)
}

func (s *Service) GetFeedItems(ctx context.Context, feed *Feed) ([]*item.Item, error) {
	return s.repo.GetFeedItems(ctx, feed)
}
