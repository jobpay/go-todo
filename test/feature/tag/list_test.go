package tag_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	tagRequest "github.com/jobpay/todo/internal/presentation/request/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

func TestListTags(t *testing.T) {
	cleanupDB()

	t.Run("正常系: タグ一覧を取得できる", func(t *testing.T) {
		// テストデータを作成
		tags := []tagRequest.StoreRequest{
			{Title: "urgent"},
			{Title: "high-priority"},
		}

		for _, tag := range tags {
			jsonBody, _ := json.Marshal(tag)
			httpClient.Post(
				testServerURL+"/api/tags",
				"application/json",
				bytes.NewBuffer(jsonBody),
			)
		}

		resp, err := httpClient.Get(testServerURL + "/api/tags")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response []tagResponse.TagResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response) != 2 {
			t.Errorf("Expected 2 tags, got %d", len(response))
		}
	})

	t.Run("正常系: タグが0件の場合空配列を返す", func(t *testing.T) {
		cleanupDB()

		resp, err := httpClient.Get(testServerURL + "/api/tags")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response []tagResponse.TagResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(response) != 0 {
			t.Errorf("Expected 0 tags, got %d", len(response))
		}
	})
}

