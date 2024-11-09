package handler

import (
	"go-test-with-db-ci/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(u usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: u}
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	todo, err := h.usecase.GetTodoByID(uint(id))
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get todo"})
	}

	return c.JSON(http.StatusOK, todo)
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

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req struct {
		Title string `json:"title" validate:"required"`
		Done  bool   `json:"done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err = h.usecase.UpdateTodo(uint(id), req.Title, req.Done)
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Todo created"})
}
