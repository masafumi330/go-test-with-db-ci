package usecase

import (
	"database/sql"
	"fmt"
	"go-test-with-db-ci/internal/repository"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// setup
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		"root",
		"password",
		"localhost", // specify host as localhost
		"3310",      // specify port as 3310
		"test_db",
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// truncate table
	db.Exec("TRUNCATE TABLE todos")

	os.Exit(m.Run())
}

func TestTodoUsecase_CreateTodo_Success(t *testing.T) {
	// Arrange
	// テストが終わったら毎回ロールバックする
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	t.Cleanup(func() {
		tx.Rollback()
	})
	repo := repository.NewTodoRepository(tx)
	usecase := NewTodoUsecase(repo)

	// Act
	err = usecase.CreateTodo("Test Todo")
	if err != nil {
		t.Errorf("failed to create todo: %v", err)
	}

	// Assert
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM todos WHERE title = ? AND done = false", "Test Todo").Scan(&count)
	if err != nil {
		t.Fatalf("failed to query todos table: %v", err)
	}
	assert.Equal(t, 1, count, "expected one record in the todos table")
}

func TestTodoUsecase_UpdateTodo_Success(t *testing.T) {
	// Arrange
	// テストが終わったら毎回ロールバックする
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	t.Cleanup(func() {
		tx.Rollback()
	})
	repo := repository.NewTodoRepository(tx)
	usecase := NewTodoUsecase(repo)
	result, err := tx.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", "Test Todo", false)
	if err != nil {
		t.Fatalf("failed to insert todo: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("failed to get last insert id: %v", err)
	}

	// Act
	err = usecase.UpdateTodo(uint(id), "Updated Todo", true)
	if err != nil {
		t.Errorf("failed to create todo: %v", err)
	}

	// Assert
	var title string
	var done bool
	err = tx.QueryRow("SELECT title, done FROM todos WHERE id = ?", id).Scan(&title, &done)
	if err != nil {
		t.Fatalf("failed to query todos table: %v", err)
	}
	assert.Equal(t, "Updated Todo", title, "expected title to be Updated Todo")
	assert.True(t, done, "expected done to be true")
}

func TestTodoUsecase_DeleteTodo_Success(t *testing.T) {
	// Arrange
	// テストが終わったら毎回ロールバックする
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	t.Cleanup(func() {
		tx.Rollback()
	})
	repo := repository.NewTodoRepository(tx)
	usecase := NewTodoUsecase(repo)
	result, err := tx.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", "Test Todo", false)
	if err != nil {
		t.Fatalf("failed to insert todo: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("failed to get last insert id: %v", err)
	}

	// Act
	err = usecase.DeleteTodo(uint(id))
	if err != nil {
		t.Errorf("failed to create todo: %v", err)
	}

	// Assert
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM todos WHERE id = ?", id).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query todos table: %v", err)
	}
	assert.Equal(t, 0, count, "expected no record in the todos table")
}
