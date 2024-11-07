package repository

import (
	"database/sql"
	"go-test-with-db-ci/internal/domain"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	_, err := r.DB.Exec("INSERT INTO todos (title) VALUES (?)", todo.Title)
	return err
}
