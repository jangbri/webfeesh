package feed

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jangbri/webfeesh-be/internal/web"
)

type Handler struct {
	service  *Service
	sync     *SyncService
	workflow *Workflow
}

func NewHandler(s *Service, sync *SyncService, w *Workflow) *Handler {
	return &Handler{
		service:  s,
		sync:     sync,
		workflow: w,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload Feed
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse json body",
		)
		return
	}

	err = h.workflow.CreateAndSync(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failure when creating new feed",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusCreated, "")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"provided id must be an integer",
		)
		return
	}

	payload := Feed{ID: int64(id)}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse json body",
		)
		return
	}

	err = h.workflow.UpdateAndSync(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failure when updating feed",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusOK, "")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"provided id must be an integer",
		)
		return
	}

	payload := Feed{ID: int64(id)}
	err = h.service.Delete(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"feed could not be deleted",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusNoContent, "")
}

func (h *Handler) GetFeedItems(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"provided id must be an integer",
		)
		return
	}

	payload := Feed{ID: int64(id)}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse json body",
		)
		return
	}

	items, err := h.service.GetFeedItems(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"feed could not be deleted",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusNoContent, items)
}
