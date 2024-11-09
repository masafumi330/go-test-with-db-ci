package main

import (
	"database/sql"
	"go-test-with-db-ci/api/handler"
	"go-test-with-db-ci/internal/repository"
	"go-test-with-db-ci/internal/usecase"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewTodoRepository(db)
	usecase := usecase.NewTodoUsecase(repo)
	todoHandler := handler.NewTodoHandler(*usecase)

	e := echo.New()
	e.GET("/todo/:id", todoHandler.GetTodo)
	e.POST("/todo", todoHandler.CreateTodo)

	e.Start(":8000")
}
