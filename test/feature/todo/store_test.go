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

func TestStoreTodo(t *testing.T) {
	cleanupDB()

	t.Run("正常系: TODOを作成できる", func(t *testing.T) {
		fixedTime := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

		reqBody := todoRequest.StoreRequest{
			Title:       "Feature Test TODO",
			Description: "This is a feature test",
			DueDate:     fixedTime,
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/todos",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		var response todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.ID <= 0 {
			t.Errorf("Expected ID > 0, got %d", response.ID)
		}

		if response.Title != reqBody.Title {
			t.Errorf("Expected Title '%s', got '%s'", reqBody.Title, response.Title)
		}

		if response.Description != reqBody.Description {
			t.Errorf("Expected Description '%s', got '%s'", reqBody.Description, response.Description)
		}

		if !response.DueDate.Truncate(time.Second).Equal(reqBody.DueDate.Truncate(time.Second)) {
			t.Errorf("Expected DueDate '%v', got '%v'", reqBody.DueDate, response.DueDate)
		}

		if response.Completed != false {
			t.Errorf("Expected Completed false, got %v", response.Completed)
		}

		var count int64
		testDB.Table("todos").Where("id = ?", response.ID).Count(&count)
		if count != 1 {
			t.Errorf("Expected 1 record in DB, got %d", count)
		}
	})

	t.Run("正常系: タグ付きTODOを作成できる", func(t *testing.T) {
		tagReqBody := map[string]string{"title": "test-tag-1"}
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

		fixedTime := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)
		reqBody := map[string]interface{}{
			"title":       "Tagged TODO",
			"description": "TODO with tags",
			"due_date":    fixedTime,
			"tag_ids":     []int{tagID},
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/todos",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		var response todoResponse.TodoResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response.Tags) != 1 {
			t.Errorf("Expected 1 tag, got %d", len(response.Tags))
		}

		if response.Tags[0].ID != tagID {
			t.Errorf("Expected tag ID %d, got %d", tagID, response.Tags[0].ID)
		}

		if response.Tags[0].Title != "test-tag-1" {
			t.Errorf("Expected tag title 'test-tag-1', got '%s'", response.Tags[0].Title)
		}
	})

	t.Run("異常系: タイトルが空の場合エラー", func(t *testing.T) {
		fixedTime := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

		reqBody := todoRequest.StoreRequest{
			Title:       "",
			Description: "No title",
			DueDate:     fixedTime,
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/todos",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
