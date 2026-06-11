package app

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/jangbri/webfeesh"
	"github.com/jangbri/webfeesh/internal/collection"
	"github.com/jangbri/webfeesh/internal/collection/feeditem"
	"github.com/jangbri/webfeesh/internal/feed"
	"github.com/jangbri/webfeesh/internal/item"
	"github.com/jangbri/webfeesh/internal/jobs"
)

type dependency struct {
	Routes *http.ServeMux
}

func BuildDependencies(db *sql.DB) *dependency {
	migrateToLatest(db)

	mux := http.NewServeMux()

	embedFrontendAssets(mux, "frontend/dist")
	setup(db, mux)

	deps := &dependency{
		Routes: mux,
	}
	return deps
}

func embedFrontendAssets(mux *http.ServeMux, path string) {
	assets, err := fs.Sub(webfeesh.Frontend, path)
	if err != nil {
		panic(err)
	}

	frontendFS := http.FileServer(http.FS(assets))
	mux.Handle("/", frontendFS)

	slog.Info("access frontend dist embedded into binary at the same port")
}

func setup(db *sql.DB, mux *http.ServeMux) {
	collectionRepo := collection.NewSQLiteRepository(db)
	feeditemRepo := feeditem.NewSQLiteRepository(db)
	feedRepo := feed.NewSQLiteRepository(db)
	itemRepo := item.NewSQLiteRepository(db)

	collectionService := collection.NewService(collectionRepo)
	feeditemService := feeditem.NewService(feeditemRepo)

	feedService := feed.NewService(feedRepo)
	syncService := feed.NewSyncService(
		feedRepo,
		itemRepo,
	)
	feedWorkflow := feed.NewWorkflow(
		feedService,
		syncService,
	)

	feedSyncJob := jobs.NewFeedSyncJob(feedWorkflow)
	feedAggJob := jobs.NewFeedItemAggJob(collectionService, feedService, itemRepo, feeditemService)
	ctx := context.Background()
	go feedSyncJob.Run(ctx)
	go feedAggJob.Run(ctx)

	// handlers
	collectionHandler := collection.NewHandler(collectionService, feeditemService)
	feedHandler := feed.NewHandler(
		feedService,
		syncService,
		feedWorkflow,
	)

	// routes
	collection.RegisterRoutes(mux, collectionHandler)
	feed.RegisterRoutes(mux, feedHandler)
}
