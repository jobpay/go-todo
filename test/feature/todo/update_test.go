package todo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	todoPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/todo"
	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

func TestUpdateTodo(t *testing.T) {
	cleanupDB()

	fixedTimeCreate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)

	createReq := todoRequest.StoreRequest{
		Title:       "Update Test TODO",
		Description: "Before update",
		DueDate:     fixedTimeCreate,
	}
	jsonBody, _ := json.Marshal(createReq)
	resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
	var created todoResponse.TodoResponse
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	t.Run("正常系: TODOを更新できる", func(t *testing.T) {
		fixedTimeUpdate := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

		updateReq := todoRequest.UpdateRequest{
			Title:       "Updated TODO",
			Description: "After update",
			Completed:   true,
			DueDate:     fixedTimeUpdate,
		}
		jsonBody, _ := json.Marshal(updateReq)

		req, _ := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/api/todos/%d", testServerURL, created.ID),
			bytes.NewBuffer(jsonBody),
		)
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.ID != created.ID {
			t.Errorf("Expected ID %d, got %d", created.ID, response.ID)
		}

		if response.Title != updateReq.Title {
			t.Errorf("Expected Title '%s', got '%s'", updateReq.Title, response.Title)
		}

		if response.Description != updateReq.Description {
			t.Errorf("Expected Description '%s', got '%s'", updateReq.Description, response.Description)
		}

		if !response.DueDate.Truncate(time.Second).Equal(updateReq.DueDate.Truncate(time.Second)) {
			t.Errorf("Expected DueDate '%v', got '%v'", updateReq.DueDate, response.DueDate)
		}

		if response.Completed != updateReq.Completed {
			t.Errorf("Expected Completed %v, got %v", updateReq.Completed, response.Completed)
		}

		var todo todoPersistence.TodoModel
		testDB.First(&todo, created.ID)
		if todo.Title != updateReq.Title {
			t.Errorf("Expected DB Title '%s', got '%s'", updateReq.Title, todo.Title)
		}
		if todo.Description != updateReq.Description {
			t.Errorf("Expected DB Description '%s', got '%s'", updateReq.Description, todo.Description)
		}
		if todo.Completed != updateReq.Completed {
			t.Errorf("Expected DB Completed %v, got %v", updateReq.Completed, todo.Completed)
		}
		if !todo.DueDate.Truncate(time.Second).Equal(updateReq.DueDate.Truncate(time.Second)) {
			t.Errorf("Expected DB DueDate '%v', got '%v'", updateReq.DueDate, todo.DueDate)
		}
	})

	t.Run("異常系: 存在しないIDでエラー", func(t *testing.T) {
		fixedTimeUpdate := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

		updateReq := todoRequest.UpdateRequest{
			Title:       "Updated TODO",
			Description: "After update",
			Completed:   true,
			DueDate:     fixedTimeUpdate,
		}
		jsonBody, _ := json.Marshal(updateReq)

		req, _ := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/api/todos/99999", testServerURL),
			bytes.NewBuffer(jsonBody),
		)
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
