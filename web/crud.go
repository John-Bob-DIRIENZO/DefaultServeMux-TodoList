package web

import (
	"demoHTTP"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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
		// Je sépare mon URL à chaque "/" et je prends le 3ème élément,
		// Le premier est une chaine vide, le second sera "delete"
		QueryId := strings.Split(request.URL.Path, "/")[2]
		id, _ := strconv.Atoi(QueryId)

		// Go ne fournit pas de méthode pour supprimer des éléments dans
		// un tableau, il faut donc le faire soi-même.
		// Je vais itérer à travers chaque élément de mon slice
		for index, todo := range h.Todos {
			// Si je tombe sur le bon élément
			if id == todo.ID {
				// Je recrée un slice qui commence par le slice originel
				// jusqu'à l'index de l'élément à supprimer
				// et qui termine par chaque élément présent à partir de
				// l'index suivant l'élément à supprimer.
				// "..." est aussi un spread operator, comme en JS
				h.Todos = append(h.Todos[:index], h.Todos[index+1:]...)
				break
			}
		}

		// Enfin, je fais ma redirection
		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}
