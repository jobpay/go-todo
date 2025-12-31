package tag_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	tagRequest "github.com/jobpay/todo/internal/presentation/request/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

func TestStoreTag(t *testing.T) {
	cleanupDB()

	t.Run("正常系: タグを作成できる", func(t *testing.T) {
		reqBody := tagRequest.StoreRequest{
			Title: "urgent",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/tags",
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

		var response tagResponse.TagResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.ID <= 0 {
			t.Errorf("Expected ID > 0, got %d", response.ID)
		}

		if response.Title != reqBody.Title {
			t.Errorf("Expected Title '%s', got '%s'", reqBody.Title, response.Title)
		}

		var count int64
		testDB.Table("tags").Where("id = ?", response.ID).Count(&count)
		if count != 1 {
			t.Errorf("Expected 1 record in DB, got %d", count)
		}
	})

	t.Run("異常系: タイトルが空の場合エラー", func(t *testing.T) {
		reqBody := tagRequest.StoreRequest{
			Title: "",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/tags",
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

	t.Run("異常系: タイトルが100文字を超える場合エラー", func(t *testing.T) {
		longTitle := ""
		for i := 0; i < 101; i++ {
			longTitle += "a"
		}
		reqBody := tagRequest.StoreRequest{
			Title: longTitle,
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/tags",
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

