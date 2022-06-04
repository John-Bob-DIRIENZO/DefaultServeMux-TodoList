package web

import (
	"demoHTTP"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetTodos() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

		todos, _ := h.Store.GetTodos()
		err := json.NewEncoder(writer).Encode(todos)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		item := demoHTTP.TodoItem{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.Store.AddTodo(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status    string `json:"status"`
			Message   string `json:"message"`
			NewTodoId int    `json:"newTodoId"`
		}{
			Status:    "success",
			Message:   "Nouveau todo inséré avec succès",
			NewTodoId: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) DeleteTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.Store.DeleteTodo(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/api", http.StatusSeeOther)
	}
}

func (h *Handler) ToggleTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.Store.ToggleTodo(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/api", http.StatusSeeOther)
	}
}
