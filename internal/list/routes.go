package list

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /feed/{id}", ping)

	mux.HandleFunc("POST /feed/{$}", ping)
	mux.HandleFunc("POST /feed/{id}", ping)

	mux.HandleFunc("DELETE /feed/{id}", ping)
}

func ping(w http.ResponseWriter, r *http.Request) {}
