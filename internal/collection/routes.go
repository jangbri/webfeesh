package collection

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /collections/{$}", h.GetAll)
	mux.HandleFunc("GET /collections/{id}", h.GetCollectionFeeds)

	mux.HandleFunc("POST /collections/{$}", h.Create)
	mux.HandleFunc("POST /collections/{id}", h.Update)

	mux.HandleFunc("DELETE /collections/{id}", h.Delete)
}
