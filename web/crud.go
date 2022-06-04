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

		// NewEncoder va me permettre de transformer un struct en JSON.
		// Je dois juste lui préciser où encoder (writter)
		// et quoi encoder (h.Todos, ma liste de todos)
		err := json.NewEncoder(writer).Encode(h.Todos)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
			return
		}

		item := demoHTTP.TodoItem{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		h.Todos = append(h.Todos, item)

		// Je crée une struct anonyme à laquelle j'ajoute
		// tout de suite un contenu que je renvoie en JSON
		err = json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Nouveau todo inséré avec succès",
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) DeleteTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Ce routeur me permet d'accéder facilement aux paramètres
		QueryId := chi.URLParam(request, "id")
		// Attention, ce sont toujours des strings...
		id, _ := strconv.Atoi(QueryId)

		for index, todo := range h.Todos {
			if id == todo.ID {
				h.Todos = append(h.Todos[:index], h.Todos[index+1:]...)
				break
			}
		}

		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}
