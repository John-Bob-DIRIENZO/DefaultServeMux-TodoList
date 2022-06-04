package web

import (
	"demoHTTP"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"strconv"
)

type TemplateData struct {
	Titre   string
	Content any
}

func (h *Handler) WebShowTodos() http.HandlerFunc {
	data := TemplateData{Titre: "Tous les todos"}

	return func(writer http.ResponseWriter, request *http.Request) {
		todos, err := h.Store.GetTodos()
		data.Content = todos

		tmpl, err := template.ParseFiles("templates/layout.html", "templates/list.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		// Je passe mes données ici
		err = tmpl.ExecuteTemplate(writer, "layout", data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) WebCreateTodoForm() http.HandlerFunc {
	data := TemplateData{Titre: "Add a todo"}

	return func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/form.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		err = tmpl.ExecuteTemplate(writer, "layout", data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) WebAddTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		_, err = h.Store.AddTodo(demoHTTP.TodoItem{Title: request.FormValue("new-todo")})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}

func (h *Handler) WebToogleTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.Store.ToggleTodo(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}

func (h *Handler) WebDeleteTodo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.Store.DeleteTodo(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}
