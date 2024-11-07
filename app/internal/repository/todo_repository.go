package repository

import (
	"database/sql"
	"go-test-with-db-ci/internal/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	_, err := r.db.Exec("INSERT INTO todos (title) VALUES (?)", todo.Title)
	return err
}
