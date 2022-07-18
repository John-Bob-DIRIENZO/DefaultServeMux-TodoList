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
	// Placer cette déclaration avant de retourner le handler
	// permet de ne créer qu'une seule fois cette struct
	// plutôt que de la créer à chaque requête
	data := TemplateData{Titre: "Tous les todos"}

	return func(writer http.ResponseWriter, request *http.Request) {
		todos, err := h.Store.GetTodos()
		data.Content = todos

		// ParseFS fonctionne exactement comme ParseFiles mais va chercher
		// dans un fileSystem donné plutôt que dans celui de l'hôte
		tmpl, err := template.ParseFS(demoHTTP.EmbedTemplates, "templates/layout.gohtml", "templates/list.gohtml")
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
		tmpl, err := template.ParseFS(demoHTTP.EmbedTemplates, "templates/layout.gohtml", "templates/form.gohtml")
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
