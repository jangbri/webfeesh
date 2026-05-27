package feed

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /collection/{$}", ping)
	mux.HandleFunc("GET /collection/{id}", ping)

	mux.HandleFunc("POST /collection/{$}", ping)
	mux.HandleFunc("POST /collection/{id}", ping)

	mux.HandleFunc("DELETE /collection/{id}", ping)
}

func ping(w http.ResponseWriter, r *http.Request) {}
