package domain

import (
	"errors"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrEmptyTitle   = errors.New("empty title")
)

type (
	ToDoID uint
	Todo   struct {
		ID    ToDoID `db:"id"`
		Title string `db:"title"`
		Done  bool   `db:"done"`
	}
)

func NewTodo(
	title string,
) (*Todo, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return &Todo{
		Title: title,
		Done:  false,
	}, nil
}
