package feed

import "context"

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

func (w *Workflow) CreateAndSync(ctx context.Context, feed *Feed) error {
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

func (w *Workflow) UpdateAndSync(ctx context.Context, feed *Feed) error {
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
