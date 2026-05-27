package list

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /lists/{$}", h.GetAll)
	mux.HandleFunc("GET /lists/{id}", h.GetListFeeds)

	mux.HandleFunc("POST /lists/{$}", h.Create)
	mux.HandleFunc("POST /lists/{id}", h.Update)

	mux.HandleFunc("DELETE /lists/{id}", h.Delete)
}
