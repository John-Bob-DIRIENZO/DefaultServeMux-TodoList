package web

import (
	"demoHTTP"
	"net/http"
)

func NewHandler(todos []demoHTTP.TodoItem) *Handler {
	handler := &Handler{
		http.NewServeMux(),
		todos,
	}

	handler.Handle("/", handler.GetTodos())
	handler.Handle("/add", handler.AddTodo())
	// Le "/" à la fin est important, sinon le routeur ne
	// comprend pas qu'il y aura des choses après "delete"
	// hors, nous voulons faire "/delete/{id}"
	handler.Handle("/delete/", handler.DeleteTodo())

	return handler
}

type Handler struct {
	*http.ServeMux
	Todos []demoHTTP.TodoItem
}
