package collection

import (
	"encoding/json"
	"encoding/xml"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/jangbri/webfeesh/internal/collection/feeditem"
	"github.com/jangbri/webfeesh/internal/web"
)

type Handler struct {
	service          *Service
	aggregateService *feeditem.Service
}

func NewHandler(s *Service, agg *feeditem.Service) *Handler {
	return &Handler{
		service:          s,
		aggregateService: agg,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	collections, err := h.service.GetAll(r.Context())
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to retrieve collections",
		)
		return
	}

	if err = web.WriteJSON(w, http.StatusOK, collections); err != nil {
		slog.Error("failed to return collections")
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload Collection
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse submitted form",
		)
		return
	}

	err = h.service.Create(r.Context(), payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to create collection",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusOK, "")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"id needs to be an integer",
		)
		return
	}

	payload := Collection{ID: int64(id)}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse collection payload",
		)
		return
	}

	err = h.service.Update(r.Context(), payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to update collection",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusOK, "")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"id needs to be an integer",
		)
		return
	}

	payload := Collection{ID: int64(id)}
	err = h.service.Delete(r.Context(), payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"DELETE_FAILED",
			"failed to delete collection",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusNoContent, "")
}

func (h *Handler) GetCollectionFeeds(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"id needs to be an integer",
		)
		return
	}

	collection := Collection{ID: int64(id)}
	feeds, err := h.service.GetCollectionFeeds(r.Context(), collection)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to find collection",
		)
		return
	}

	if err = web.WriteJSON(w, http.StatusOK, feeds); err != nil {
		slog.Error("failed to write collection's feeds")
	}
}

func (h *Handler) GetAggregateRSS(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	items, err := h.aggregateService.GetLatest(r.Context(), int64(id))
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to retrieve rss items",
		)
		return
	}

	feed := feeditem.RSS{
		Version: "2.0",
		Channel: feeditem.Channel{
			Title:       r.URL.String(),
			Link:        r.URL.String(),
			Description: "Latest Aggregated Feed Items",
			Items:       items,
		},
	}

	output, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to marshal items into a proper rss feed",
		)
		return
	}

	rss := xml.Header + string(output)

	w.Header().Set("Content-Type", "application/rss+xml")
	_, err = w.Write([]byte(rss))
}
