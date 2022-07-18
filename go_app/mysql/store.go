package database

import (
	"database/sql"
	"demoHTTP"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewTodoStore(db),
	}
}

type Store struct {
	demoHTTP.TodoStoreInterface
}
