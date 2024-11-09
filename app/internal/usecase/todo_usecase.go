package usecase

import (
	"errors"
	"go-test-with-db-ci/internal/domain"
	"go-test-with-db-ci/internal/repository"
)

var ErrTodoNotFound = errors.New("todo not found")

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: *repo}
}

func (u *TodoUsecase) GetTodoByID(id uint) (*domain.Todo, error) {
	todo, err := u.repo.GetByID(domain.ToDoID(id))
	if err != nil {
		if errors.Is(err, domain.ErrTodoNotFound) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (u *TodoUsecase) CreateTodo(title string) error {
	todo, err := domain.NewTodo(title)
	if err != nil {
		return err
	}
	return u.repo.Create(todo)
}
