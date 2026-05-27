package list

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /list/{$}", h.GetAll)
	mux.HandleFunc("GET /list/{id}", h.GetListFeeds)

	mux.HandleFunc("POST /list/{$}", h.Create)
	mux.HandleFunc("POST /list/{id}", h.Update)

	mux.HandleFunc("DELETE /list/{id}", h.Delete)
}
