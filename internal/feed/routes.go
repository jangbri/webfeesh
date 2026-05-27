package feed

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /feeds/{$}", h.Create)
	mux.HandleFunc("POST /feeds/{id}", h.Update)

	mux.HandleFunc("DELETE /feeds/{id}", h.Delete)

	mux.HandleFunc("GET /feeds/{id}", h.GetFeedItems)
}
