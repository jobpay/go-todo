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

func TestShowTodo(t *testing.T) {
	cleanupDB()

	fixedTime := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	reqBody := todoRequest.StoreRequest{
		Title:       "Show Test TODO",
		Description: "For show test",
		DueDate:     fixedTime,
	}
	jsonBody, _ := json.Marshal(reqBody)
	resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
	var created todoResponse.TodoResponse
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	t.Run("正常系: IDでTODOを取得できる", func(t *testing.T) {
		resp, err := httpClient.Get(fmt.Sprintf("%s/api/todos/%d", testServerURL, created.ID))
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
	})

	t.Run("正常系: タグ付きTODOを取得できる", func(t *testing.T) {
		tagReqBody := map[string]string{"title": "show-test-tag"}
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
			"title":       "Tagged Show TODO",
			"description": "TODO with tags for show",
			"due_date":    fixedTime,
			"tag_ids":     []int{tagID},
		}
		jsonBody, _ := json.Marshal(reqBody)
		resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
		var created todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		resp, err := httpClient.Get(fmt.Sprintf("%s/api/todos/%d", testServerURL, created.ID))
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

		if len(response.Tags) != 1 {
			t.Errorf("Expected 1 tag, got %d", len(response.Tags))
		}

		if response.Tags[0].ID != tagID {
			t.Errorf("Expected tag ID %d, got %d", tagID, response.Tags[0].ID)
		}

		if response.Tags[0].Title != "show-test-tag" {
			t.Errorf("Expected tag title 'show-test-tag', got '%s'", response.Tags[0].Title)
		}
	})

	t.Run("異常系: 存在しないIDでエラー", func(t *testing.T) {
		resp, err := httpClient.Get(fmt.Sprintf("%s/api/todos/99999", testServerURL))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
