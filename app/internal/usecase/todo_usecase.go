package usecase

import (
	"go-test-with-db-ci/internal/domain"
	"go-test-with-db-ci/internal/repository"
)

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: *repo}
}

func (u *TodoUsecase) CreateTodo(title string) error {
	todo, err := domain.NewTodo(title)
	if err != nil {
		return err
	}
	return u.repo.Create(todo)
}
