package handler

import (
	"refactoring/internal/adapters/db"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	storage db.UserStore
}

func New(storage db.UserStore) handler {
	return handler{
		storage: storage,
	}
}

func (h *handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", h.getTimeNow)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", h.searchUsers)
				r.Post("/", h.createUser)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.getUser)
					r.Patch("/", h.updateUser)
					r.Delete("/", h.deleteUser)
				})
			})
		})
	})
	return r
}
