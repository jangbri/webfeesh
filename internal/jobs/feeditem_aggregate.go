package jobs

import (
	"bytes"
	"context"
	"html/template"
	"log/slog"
	"strings"
	"time"

	"github.com/jangbri/webfeesh/internal/collection"
	"github.com/jangbri/webfeesh/internal/collection/feeditem"
	"github.com/jangbri/webfeesh/internal/feed"
	"github.com/jangbri/webfeesh/internal/item"
)

type FeedItemAggJob struct {
	collectionService *collection.Service
	feedService       *feed.Service
	itemRepo          *item.SQLiteRepository
	feeditemService   *feeditem.Service
}

func NewFeedItemAggJob(
	c *collection.Service,
	f *feed.Service,
	i *item.SQLiteRepository,
	fi *feeditem.Service,
) *FeedItemAggJob {
	return &FeedItemAggJob{
		collectionService: c,
		feedService:       f,
		itemRepo:          i,
		feeditemService:   fi,
	}
}

func (j *FeedItemAggJob) Run(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			slog.Info("running feed items aggregate job")

			err := j.aggregate(ctx)
			if err != nil {
				slog.Error("feed sync failed")
			}
		}
	}
}

// inefficient time complexity: O(n^3)
// but it runs infrequently
// once every 12 hours which is more than enough leeway
func (j *FeedItemAggJob) aggregate(ctx context.Context) error {
	collections, err := j.collectionService.GetAll(ctx)
	if err != nil {
		slog.Error("failed to retrieve collections to aggregate")
	}

	for _, collection := range collections {
		updatedFeedItems := make(map[string][]*item.Item)

		latestPrevRetrieval := j.latestPrevRetrieval(ctx, collection.ID)

		feeds, err := j.collectionService.GetCollectionFeeds(ctx, collection)
		if err != nil {
			slog.Error("failed to retrieve collection's feeds to aggregate")
		}

		for _, feed := range feeds {
			items := j.itemsSinceTimestamp(ctx, latestPrevRetrieval, feed)
			if len(items) > 0 {
				updatedFeedItems[feed.Title] = items
			}
		}

		if len(updatedFeedItems) == 0 {
			continue
		}

		// create feeditem row entry per collection, the important part is the desc
		funcMap := template.FuncMap{
			"truncate": truncate,
		}
		tmpl := template.Must(
			template.New("content").Funcs(funcMap).Parse(`
			{{ range $feedName, $items := . }}
				<h2>{{ $feedName }}</h2>
				<ul>
					{{ range $items }}
						<li> <a href="{{ .Link }}">{{ truncate 17 .Title }}</a> {{ .DateUpdated }}</li>
					{{ end }}
				</ul>
			{{ end }}
		`),
		)
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, updatedFeedItems)

		var b strings.Builder
		b.WriteString("<![CDATA[")
		b.WriteString(buf.String())
		b.WriteString("]]")

		entry := feeditem.CollectionFeedItem{
			ID:           -1,
			CollectionID: collection.ID,
			Title:        collection.Name,
			Description:  b.String(),
			TimeCreated:  time.Now().UTC(),
		}

		err = j.feeditemService.Create(ctx, &entry)
		if err != nil {
			slog.Error("issue encountered when attempting to create collection feed item")
			return err
		}
	}

	return nil
}

func (j *FeedItemAggJob) latestPrevRetrieval(ctx context.Context, cID int64) time.Time {
	feeditems, err := j.feeditemService.GetLatest(ctx, cID)
	if err != nil || len(feeditems) == 0 {
		return time.Unix(0, 0).UTC()
	}
	return feeditems[0].TimeCreated.UTC()
}

func (j *FeedItemAggJob) itemsSinceTimestamp(
	ctx context.Context,
	t time.Time,
	f *feed.Feed,
) []*item.Item {
	items, err := j.feedService.GetFeedItems(ctx, f)
	if err != nil {
		return []*item.Item{}
	}

	desiredItems := make([]*item.Item, 0, len(items))
	for _, item := range items {
		if t.Before(item.DateUpdated) {
			desiredItems = append(desiredItems, item)
		}
	}

	return desiredItems
}

func truncate(n int, s string) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
