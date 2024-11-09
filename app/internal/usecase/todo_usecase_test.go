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
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
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
