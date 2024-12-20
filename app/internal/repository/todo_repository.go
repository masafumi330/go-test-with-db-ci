package repository

import (
	"database/sql"
	"errors"
	"go-test-with-db-ci/internal/domain"
)

type CustomDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TodoRepository struct {
	db CustomDB
}

func NewTodoRepository(db CustomDB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetByID(id domain.ToDoID) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.QueryRow("SELECT id, title, done FROM todos WHERE id = ? LIMIT 1", id).Scan(&todo.ID, &todo.Title, &todo.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrTodoNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	_, err := r.db.Exec("INSERT INTO todos (title) VALUES (?)", todo.Title)
	return err
}

func (r *TodoRepository) Update(todo *domain.Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", todo.Title, todo.Done, todo.ID)
	return err
}

func (r *TodoRepository) Delete(id domain.ToDoID) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}
