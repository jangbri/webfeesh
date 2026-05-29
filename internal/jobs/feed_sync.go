package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/jangbri/webfeesh-be/internal/feed"
)

type FeedSyncJob struct {
	workflow *feed.Workflow
}

func NewFeedSyncJob(workflow *feed.Workflow) *FeedSyncJob {
	return &FeedSyncJob{
		workflow: workflow,
	}
}

func (j *FeedSyncJob) Run(ctx context.Context) {
	// run first time the app starts, then 30 minute intervals
	err := j.workflow.SyncAllFeeds(ctx)
	if err != nil {
		slog.Error("feed sync failed")
	}

	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			slog.Info("running feed sync job")

			err := j.workflow.SyncAllFeeds(ctx)
			if err != nil {
				slog.Error("feed sync failed")
			}
		}
	}
}
