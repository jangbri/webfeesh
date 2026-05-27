package app

import (
	"database/sql"
	"net/http"

	"github.com/jangbri/webfeesh-be/internal/collection"
	"github.com/jangbri/webfeesh-be/internal/feed"
	"github.com/jangbri/webfeesh-be/internal/item"
)

type dependency struct {
	Routes *http.ServeMux
}

func BuildDependencies(db *sql.DB) *dependency {
	migrateToLatest(db)

	mux := http.NewServeMux()

	// setup models, views, and controllers
	buildMVC(db, mux)

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/collection/{$}", http.StatusFound)
	})

	deps := &dependency{
		Routes: mux,
	}
	return deps
}

func buildMVC(db *sql.DB, mux *http.ServeMux) {
	collectionRepo := collection.NewSQLiteRepository(db)
	feedRepo := feed.NewSQLiteRepository(db)
	itemRepo := item.NewSQLiteRepository(db)

	collectionService := collection.NewService(collectionRepo)

	feedService := feed.NewService(feedRepo)
	syncService := feed.NewSyncService(
		feedRepo,
		itemRepo,
	)
	feedWorkflow := feed.NewWorkflow(
		feedService,
		syncService,
	)

	// handlers
	collectionHandler := collection.NewHandler(collectionService)
	feedHandler := feed.NewHandler(
		feedService,
		syncService,
		feedWorkflow,
	)

	// routes
	collection.RegisterRoutes(mux, collectionHandler)
	feed.RegisterRoutes(mux, feedHandler)
}
