package collection

import "github.com/go-chi/chi"

func RegisterRoutes(h *Handler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetCollectionFeeds)

	r.Post("/", h.Create)
	r.Post("/{id}", h.Update)

	r.Delete("/{id}", h.Delete)
	r.Get("/{id}/rss", h.GetAggregateRSS)

	return r
}
