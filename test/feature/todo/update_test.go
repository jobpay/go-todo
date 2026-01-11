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

	fixedTimeCreate := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

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

	t.Run("正常系: タグを追加できる", func(t *testing.T) {
		cleanupDB()

		// Todoを作成（タグなし）
		createReq := map[string]interface{}{
			"title":       "TODO without tags",
			"description": "No tags initially",
			"due_date":    fixedTimeCreate,
		}
		jsonBody, _ := json.Marshal(createReq)
		resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
		var created todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// タグを作成
		tagReqBody := map[string]string{"title": "update-test-tag"}
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

		// タグを追加
		updateReq := map[string]interface{}{
			"title":       "TODO with tags",
			"description": "Tags added",
			"completed":   false,
			"due_date":    fixedTimeCreate,
			"tag_ids":     []int{tagID},
		}
		jsonBody, _ = json.Marshal(updateReq)
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

		if len(response.Tags) != 1 {
			t.Errorf("Expected 1 tag, got %d", len(response.Tags))
		}

		if response.Tags[0].ID != tagID {
			t.Errorf("Expected tag ID %d, got %d", tagID, response.Tags[0].ID)
		}
	})

	t.Run("正常系: タグを変更できる", func(t *testing.T) {
		cleanupDB()

		// タグを2つ作成
		tag1ReqBody := map[string]string{"title": "tag1"}
		tag1Json, _ := json.Marshal(tag1ReqBody)
		tag1Resp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(tag1Json),
		)
		var tag1Response map[string]interface{}
		json.NewDecoder(tag1Resp.Body).Decode(&tag1Response)
		tag1Resp.Body.Close()
		tag1ID := int(tag1Response["id"].(float64))

		tag2ReqBody := map[string]string{"title": "tag2"}
		tag2Json, _ := json.Marshal(tag2ReqBody)
		tag2Resp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(tag2Json),
		)
		var tag2Response map[string]interface{}
		json.NewDecoder(tag2Resp.Body).Decode(&tag2Response)
		tag2Resp.Body.Close()
		tag2ID := int(tag2Response["id"].(float64))

		// Tag1付きTodoを作成
		createReq := map[string]interface{}{
			"title":       "TODO with tag1",
			"description": "Has tag1",
			"due_date":    fixedTimeCreate,
			"tag_ids":     []int{tag1ID},
		}
		jsonBody, _ := json.Marshal(createReq)
		resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
		var created todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// Tag2に変更
		updateReq := map[string]interface{}{
			"title":       "TODO with tag2",
			"description": "Changed to tag2",
			"completed":   false,
			"due_date":    fixedTimeCreate,
			"tag_ids":     []int{tag2ID},
		}
		jsonBody, _ = json.Marshal(updateReq)
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

		var response todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&response)

		if len(response.Tags) != 1 {
			t.Errorf("Expected 1 tag, got %d", len(response.Tags))
		}

		if response.Tags[0].ID != tag2ID {
			t.Errorf("Expected tag ID %d, got %d", tag2ID, response.Tags[0].ID)
		}

		if response.Tags[0].Title != "tag2" {
			t.Errorf("Expected tag title 'tag2', got '%s'", response.Tags[0].Title)
		}
	})

	t.Run("正常系: タグを削除できる", func(t *testing.T) {
		cleanupDB()

		// タグを作成
		tagReqBody := map[string]string{"title": "tag-to-remove"}
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

		// タグ付きTodoを作成
		createReq := map[string]interface{}{
			"title":       "TODO with tag",
			"description": "Has tag",
			"due_date":    fixedTimeCreate,
			"tag_ids":     []int{tagID},
		}
		jsonBody, _ := json.Marshal(createReq)
		resp, _ := httpClient.Post(testServerURL+"/api/todos", "application/json", bytes.NewBuffer(jsonBody))
		var created todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// タグを削除
		updateReq := map[string]interface{}{
			"title":       "TODO without tag",
			"description": "Tag removed",
			"completed":   false,
			"due_date":    fixedTimeCreate,
			"tag_ids":     []int{},
		}
		jsonBody, _ = json.Marshal(updateReq)
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

		var response todoResponse.TodoResponse
		json.NewDecoder(resp.Body).Decode(&response)

		if len(response.Tags) != 0 {
			t.Errorf("Expected 0 tags, got %d", len(response.Tags))
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
