package feed

import (
	"github.com/go-chi/chi"
)

func RegisterRoutes(h *Handler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Post("/{id}", h.Update)

	r.Delete("/{id}", h.Delete)

	r.Get("/{id}", h.GetFeedItems)

	return r
}
