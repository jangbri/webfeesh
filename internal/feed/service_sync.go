package feed

import (
	"context"
	"log/slog"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/jangbri/webfeesh/internal/item"
)

type SyncService struct {
	feedRepo Repository
	itemRepo item.Repository
	parser   *gofeed.Parser
}

func NewSyncService(fp Repository, ip item.Repository) *SyncService {
	return &SyncService{
		feedRepo: fp,
		itemRepo: ip,
		parser:   gofeed.NewParser(),
	}
}

func (s *SyncService) SyncFeed(ctx context.Context, feed Feed) error {
	parsed, err := s.parser.ParseURLWithContext(feed.Link, ctx)
	if err != nil {
		slog.Error("failed to parse feed link")
		return err
	}

	items := []*item.Item{}
	for _, i := range parsed.Items {

		// INFO: cleanup a bunch of potentially missing values

		if len(i.Link) == 0 {
			i.Link = feed.Link
		}

		if len(i.GUID) == 0 {
			i.GUID = i.Link + i.Title + i.Description
		}

		var dateUpdated time.Time
		if i.UpdatedParsed != nil {
			dateUpdated = i.UpdatedParsed.UTC()
		} else if i.PublishedParsed != nil {
			dateUpdated = i.PublishedParsed.UTC()
		} else {
			dateUpdated = time.Now().UTC()
		}

		item := item.Item{
			ID:          -1,
			FeedID:      feed.ID,
			GUID:        i.GUID,
			Link:        i.Link,
			Title:       i.Title,
			DateFetched: time.Now().UTC(),
			DateUpdated: dateUpdated,
		}

		items = append(items, &item)
	}

	err = s.itemRepo.UpsertMany(ctx, items)
	if err != nil {
		slog.Error("failed to upsert multiple items")
		return err
	}

	return nil
}
