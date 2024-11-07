package domain

import (
	"errors"
)

var ErrEmptyTitle = errors.New("empty title")

type Todo struct {
	ID    uint
	Title string
	Done  bool
}

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
