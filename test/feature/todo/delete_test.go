package todo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

func TestDeleteTodo(t *testing.T) {
	cleanupDB()

	fixedTime := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	createReq := todoRequest.StoreRequest{
		Title:       "Delete Test TODO",
		Description: "Will be deleted",
		DueDate:     fixedTime,
	}
	jsonBody, _ := json.Marshal(createReq)
	resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
	var created todoResponse.TodoResponse
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	t.Run("正常系: TODOを削除できる", func(t *testing.T) {
		req, _ := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/api/todos/%d", testServerURL, created.ID),
			nil,
		)

		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
		}

		var count int64
		testDB.Table("todos").Where("id = ?", created.ID).Count(&count)
		if count != 0 {
			t.Errorf("Expected 0 records in DB, got %d", count)
		}
	})

	t.Run("異常系: 存在しないIDでエラー", func(t *testing.T) {
		req, _ := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/api/todos/99999", testServerURL),
			nil,
		)

		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
