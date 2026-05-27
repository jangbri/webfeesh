package app

import (
	"database/sql"
	"net/http"
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
}
