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

	handler.Get("/", handler.WebShowTodos())
	handler.Get("/create-todo", handler.WebCreateTodoForm())
	handler.Post("/add-todo", handler.WebAddTodo())
	handler.Get("/toggle-todo/{id}", handler.WebToogleTodo())
	handler.Get("/delete-todo/{id}", handler.WebDeleteTodo())

	handler.Route("/api", func(r chi.Router) {
		r.Get("/", handler.GetTodos())
		r.Post("/", handler.AddTodo())
		r.Delete("/{id}", handler.DeleteTodo())
		r.Patch("/{id}", handler.ToggleTodo())
	})

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
