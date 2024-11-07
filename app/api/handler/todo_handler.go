package handler

import (
	"go-test-with-db-ci/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(u usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: u}
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var req struct {
		Title string `json:"title" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.usecase.CreateTodo(req.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create todo"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Todo created"})
}
