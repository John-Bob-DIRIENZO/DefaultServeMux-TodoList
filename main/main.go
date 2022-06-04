package main

import (
	"demoHTTP"
	"demoHTTP/web"
	"fmt"
	"net/http"
)

var Todos = []demoHTTP.TodoItem{
	{
		ID:        1,
		Title:     "Faire la vaisselle",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Laver la voiture",
		Completed: false,
	},
	{
		ID:        3,
		Title:     "Apprendre Go",
		Completed: true,
	},
}

func main() {
	mux := web.NewHandler(Todos)

	err := http.ListenAndServe(":8097", mux)
	if err != nil {
		_ = fmt.Errorf("impossible de lancer le serveur : %w", err)
		return
	}
}
