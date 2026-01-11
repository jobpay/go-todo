package todo_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

func TestListTodos(t *testing.T) {
	cleanupDB()

	fixedTime1 := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)
	fixedTime2 := time.Date(2027, 1, 15, 12, 0, 0, 0, time.UTC)

	testTodos := []todoRequest.StoreRequest{
		{
			Title:       "List Test TODO 1",
			Description: "First todo",
			DueDate:     fixedTime1,
		},
		{
			Title:       "List Test TODO 2",
			Description: "Second todo",
			DueDate:     fixedTime2,
		},
	}

	for _, todo := range testTodos {
		jsonBody, _ := json.Marshal(todo)
		httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
	}

	t.Run("正常系: TODO一覧を取得できる", func(t *testing.T) {
		resp, err := httpClient.Get(testServerURL + "/api/todos")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response []todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response) != 2 {
			t.Fatalf("Expected 2 TODOs, got %d", len(response))
		}

		expectedTodos := map[string]struct {
			Description string
			DueDate     time.Time
			Completed   bool
		}{
			"List Test TODO 1": {
				Description: "First todo",
				DueDate:     testTodos[0].DueDate,
				Completed:   false,
			},
			"List Test TODO 2": {
				Description: "Second todo",
				DueDate:     testTodos[1].DueDate,
				Completed:   false,
			},
		}

		for _, todo := range response {
			expected, exists := expectedTodos[todo.Title]
			if !exists {
				t.Errorf("Unexpected TODO title: %s", todo.Title)
				continue
			}

			if todo.Description != expected.Description {
				t.Errorf("TODO '%s': Expected Description '%s', got '%s'",
					todo.Title, expected.Description, todo.Description)
			}

			if !todo.DueDate.Truncate(time.Second).Equal(expected.DueDate.Truncate(time.Second)) {
				t.Errorf("TODO '%s': Expected DueDate '%v', got '%v'",
					todo.Title, expected.DueDate, todo.DueDate)
			}

			if todo.Completed != expected.Completed {
				t.Errorf("TODO '%s': Expected Completed %v, got %v",
					todo.Title, expected.Completed, todo.Completed)
			}

			if todo.ID <= 0 {
				t.Errorf("TODO '%s': Expected ID > 0, got %d", todo.Title, todo.ID)
			}

			delete(expectedTodos, todo.Title)
		}

		if len(expectedTodos) > 0 {
			for title := range expectedTodos {
				t.Errorf("Expected TODO '%s' not found in response", title)
			}
		}
	})

	t.Run("正常系: タグ付きTODO一覧を取得できる", func(t *testing.T) {
		cleanupDB()

		tagReqBody := map[string]string{"title": "list-test-tag"}
		tagJson, _ := json.Marshal(tagReqBody)
		tagResp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(tagJson),
		)
		var tagResponse map[string]interface{}
		json.NewDecoder(tagResp.Body).Decode(&tagResponse)
		tagResp.Body.Close()
		tagID := int(tagResponse["id"].(float64))

		reqBody := map[string]interface{}{
			"title":       "Tagged List TODO",
			"description": "TODO with tags for list",
			"due_date":    fixedTime1,
			"tag_ids":     []int{tagID},
		}
		jsonBody, _ := json.Marshal(reqBody)
		httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))

		resp, err := httpClient.Get(testServerURL + "/api/todos")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response []todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response) != 1 {
			t.Fatalf("Expected 1 TODO, got %d", len(response))
		}

		if len(response[0].Tags) != 1 {
			t.Errorf("Expected 1 tag, got %d", len(response[0].Tags))
		}

		if response[0].Tags[0].ID != tagID {
			t.Errorf("Expected tag ID %d, got %d", tagID, response[0].Tags[0].ID)
		}

		if response[0].Tags[0].Title != "list-test-tag" {
			t.Errorf("Expected tag title 'list-test-tag', got '%s'", response[0].Tags[0].Title)
		}
	})

	t.Run("正常系: 空のリストを取得できる", func(t *testing.T) {
		cleanupDB()

		resp, err := httpClient.Get(testServerURL + "/api/todos")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response []todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response) != 0 {
			t.Errorf("Expected 0 TODOs, got %d", len(response))
		}
	})
}
