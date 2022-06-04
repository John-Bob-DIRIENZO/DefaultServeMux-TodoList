package web

import (
	database "demoHTTP/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", handler.GetTodos())
	handler.Post("/", handler.AddTodo())
	handler.Delete("/{id}", handler.DeleteTodo())
	handler.Patch("/{id}", handler.ToggleTodo())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
