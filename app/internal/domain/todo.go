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

// TodoOption は、Todo コンストラクタに渡すオプションを表す型
type TodoOption func(*Todo) error

// WithID は、ID を設定するオプション
func WithID(id ToDoID) TodoOption {
	return func(todo *Todo) error {
		todo.ID = id
		return nil
	}
}

// NewTodo はファンクショナルオプションを受け取るコンストラクタ
func NewTodo(
	title string,
	done bool,
	opts ...TodoOption,
) (*Todo, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	todo := &Todo{
		Title: title,
		Done:  done,
	}

	// オプションを適用
	for _, opt := range opts {
		if err := opt(todo); err != nil {
			return nil, err
		}
	}
	return todo, nil
}
