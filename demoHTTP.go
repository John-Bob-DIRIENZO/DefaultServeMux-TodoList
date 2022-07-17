package demoHTTP

import "embed"

//go:embed templates/*
var EmbedTemplates embed.FS

type TodoItem struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoStoreInterface interface {
	GetTodos() ([]TodoItem, error)
	AddTodo(item TodoItem) (int, error)
	DeleteTodo(id int) error
	ToggleTodo(id int) error
}
