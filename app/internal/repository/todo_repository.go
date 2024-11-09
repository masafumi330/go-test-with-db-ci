package repository

import (
	"database/sql"
	"errors"
	"go-test-with-db-ci/internal/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
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
