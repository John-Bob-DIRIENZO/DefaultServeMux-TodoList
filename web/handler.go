package web

import (
	"demoHTTP"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(todos []demoHTTP.TodoItem) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		todos,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", handler.GetTodos())
	handler.Post("/", handler.AddTodo())
	handler.Delete("/{id}", handler.DeleteTodo())

	return handler
}

type Handler struct {
	*chi.Mux
	Todos []demoHTTP.TodoItem
}
