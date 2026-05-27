package list

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

func (s *Service) GetAll(ctx context.Context) ([]*List, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, list *List) error {
	if list.Name == "" {
		return errors.New("need to provide a list name")
	}
	return s.repo.Create(ctx, list)
}

func (s *Service) Update(ctx context.Context, list *List) error {
	if list.Name == "" {
		return errors.New("need to provide a list name")
	}
	return s.repo.Update(ctx, list)
}

func (s *Service) Delete(ctx context.Context, list *List) error {
	return s.repo.Delete(ctx, list)
}

func (s *Service) GetListFeeds(ctx context.Context, list *List) ([]*feed.Feed, error) {
	return s.repo.GetListFeeds(ctx, list)
}
