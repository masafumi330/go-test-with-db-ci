package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	// db, err := sql.Open("mysql", "user:password@tcp(db:3306)/app_db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// repo := repository.NewTodoRepository(db)
	// usecase := usecase.NewTodoUsecase(repo)
	// todoHandler := handler.NewTodoHandler(*usecase)

	e := echo.New()
	// e.POST("/todo", todoHandler.CreateTodo)
	e.GET("/todo", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Start(":8000")
}
