package collection

import (
	"context"
	"errors"

	"github.com/jangbri/webfeesh-be/internal/feed"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*Collection, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, collection *Collection) error {
	if collection.Name == "" {
		return errors.New("need to provide a collection name")
	}
	return s.repo.Create(ctx, collection)
}

func (s *Service) Update(ctx context.Context, collection *Collection) error {
	if collection.Name == "" {
		return errors.New("need to provide a collection name")
	}
	return s.repo.Update(ctx, collection)
}

func (s *Service) Delete(ctx context.Context, collection *Collection) error {
	return s.repo.Delete(ctx, collection)
}

func (s *Service) GetCollectionFeeds(
	ctx context.Context,
	collection *Collection,
) ([]*feed.Feed, error) {
	return s.repo.GetCollectionFeeds(ctx, collection)
}
