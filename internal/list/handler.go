package list

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jangbri/webfeesh-be/internal/web"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	lists, err := h.service.GetAll(r.Context())
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to retrieve lists",
		)
		return
	}

	if err = web.WriteJSON(w, http.StatusOK, lists); err != nil {
		slog.Error("failed to return lists")
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload List
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse submitted form",
		)
		return
	}

	err = h.service.Create(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to create list",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusOK, "")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"id needs to be an integer",
		)
		return
	}

	payload := List{ID: int64(id)}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse list payload",
		)
		return
	}

	err = h.service.Update(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to update list",
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
			"id needs to be an integer",
		)
		return
	}

	payload := List{ID: int64(id)}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"BAD_REQUEST",
			"failed to parse list payload",
		)
		return
	}

	err = h.service.Delete(r.Context(), &payload)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"DELETE_FAILED",
			"failed to delete list",
		)
		return
	}

	_ = web.WriteJSON(w, http.StatusNoContent, "")
}

func (h *Handler) GetListFeeds(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.WriteError(
			w, http.StatusBadRequest,
			"INVALID_ID",
			"id needs to be an integer",
		)
		return
	}

	list := List{ID: int64(id)}
	feeds, err := h.service.GetListFeeds(r.Context(), &list)
	if err != nil {
		web.WriteError(
			w, http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"failed to find list",
		)
		return
	}

	if err = web.WriteJSON(w, http.StatusOK, feeds); err != nil {
		slog.Error("failed to write list's feeds")
	}
}
