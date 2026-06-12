package app

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/jangbri/webfeesh"
	"github.com/jangbri/webfeesh/internal/collection"
	"github.com/jangbri/webfeesh/internal/collection/feeditem"
	"github.com/jangbri/webfeesh/internal/feed"
	"github.com/jangbri/webfeesh/internal/item"
	"github.com/jangbri/webfeesh/internal/jobs"
)

type dependency struct {
	Routes *chi.Mux
}

func BuildDependencies(db *sql.DB) *dependency {
	migrateToLatest(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	setup(db, r)
	embedFrontendAssets(r, "frontend/dist")

	deps := &dependency{
		Routes: r,
	}
	return deps
}

func embedFrontendAssets(r *chi.Mux, path string) {
	assets, err := fs.Sub(webfeesh.Frontend, path)
	if err != nil {
		panic(err)
	}

	frontendFS := http.FileServer(http.FS(assets))
	r.Handle("/*", frontendFS)

	slog.Info("access frontend dist embedded into binary at the same port")
}

func setup(db *sql.DB, r *chi.Mux) {
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
	r.Route("/api", func(api chi.Router) {
		api.Mount("/collections", collection.RegisterRoutes(collectionHandler))
		api.Mount("/feeds", feed.RegisterRoutes(feedHandler))
	})
}
