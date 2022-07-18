package database

import (
	"database/sql"
	"demoHTTP"
)

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{
		db,
	}
}

type TodoStore struct {
	*sql.DB
}

func (t *TodoStore) GetTodos() ([]demoHTTP.TodoItem, error) {
	var todos []demoHTTP.TodoItem

	rows, err := t.Query("SELECT id, title, completed FROM Todos")
	if err != nil {
		return []demoHTTP.TodoItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo demoHTTP.TodoItem
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return []demoHTTP.TodoItem{}, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return []demoHTTP.TodoItem{}, err
	}

	return todos, nil
}

func (t *TodoStore) AddTodo(item demoHTTP.TodoItem) (int, error) {
	res, err := t.DB.Exec("INSERT INTO Todos (title, completed) VALUES (?, ?)", item.Title, item.Completed)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *TodoStore) DeleteTodo(id int) error {
	_, err := t.DB.Exec("DELETE FROM Todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoStore) ToggleTodo(id int) error {
	_, err := t.DB.Exec("UPDATE Todos SET `completed` = IF (`completed`, 0, 1) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
