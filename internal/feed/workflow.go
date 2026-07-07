package feed

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

type Workflow struct {
	feeds *Service
	sync  *SyncService
}

func NewWorkflow(feeds *Service, sync *SyncService) *Workflow {
	return &Workflow{
		feeds: feeds,
		sync:  sync,
	}
}

func (w *Workflow) CreateAndSync(ctx context.Context, feed Feed) error {
	feed, err := w.feeds.Create(ctx, feed)
	if err != nil {
		return err
	}

	err = w.sync.SyncFeed(ctx, feed)
	if err != nil {
		return err
	}

	return nil
}

func (w *Workflow) UpdateAndSync(ctx context.Context, feed Feed) error {
	feed, err := w.feeds.Update(ctx, feed)
	if err != nil {
		return err
	}

	err = w.sync.SyncFeed(ctx, feed)
	if err != nil {
		return err
	}

	return nil
}

func (w *Workflow) SyncAllFeeds(ctx context.Context) error {
	feeds, err := w.feeds.GetAllFeeds(ctx)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)

	g.SetLimit(20)

	for _, feed := range feeds {
		g.Go(func() error {
			err := w.sync.SyncFeed(ctx, feed)
			if err != nil {
				slog.Error("failed to parse feed", "feed", feed.Link, "err", err)
			}
			return nil
		})
	}

	return g.Wait()
}
