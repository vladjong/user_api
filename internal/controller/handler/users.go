package handler

import (
	"net/http"
	"refactoring/internal/controller/handler/dto"
	"refactoring/internal/entities"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *handler) searchUsers(w http.ResponseWriter, r *http.Request) {
	list, err := h.storage.SearchUsers()
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}
	render.JSON(w, r, list)
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	request := dto.CreateUserRequest{}
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}
	u := entities.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}
	id, err := h.storage.CreateUser(u)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h *handler) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(id); err != nil {
		render.Render(w, r, ErrInvalidRequest(ErrIncorrectUserId, http.StatusBadRequest))
		return
	}
	user, err := h.storage.GetUser(id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	render.JSON(w, r, user)
}

func (h *handler) updateUser(w http.ResponseWriter, r *http.Request) {
	request := dto.UpdateUserRequest{}
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}
	id := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(id); err != nil {
		render.Render(w, r, ErrInvalidRequest(ErrIncorrectUserId, http.StatusBadRequest))
		return
	}
	if err := h.storage.UpdateUser(id, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"status": "Ok",
	})
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(id); err != nil {
		render.Render(w, r, ErrInvalidRequest(ErrIncorrectUserId, http.StatusBadRequest))
		return
	}
	if err := h.storage.DeleteUser(id); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"status": "Ok",
	})
}
