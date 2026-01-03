package tag_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	tagRequest "github.com/jobpay/todo/internal/presentation/request/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

func TestShowTag(t *testing.T) {
	cleanupDB()

	t.Run("正常系: タグ詳細を取得できる", func(t *testing.T) {
		// テストデータを作成
		reqBody := tagRequest.StoreRequest{
			Title: "urgent",
		}
		jsonBody, _ := json.Marshal(reqBody)

		createResp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		defer createResp.Body.Close()

		var created tagResponse.TagResponse
		json.NewDecoder(createResp.Body).Decode(&created)

		// 詳細を取得
		resp, err := httpClient.Get(fmt.Sprintf("%s/api/tags/%d", testServerURL, created.ID))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response tagResponse.TagResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.ID != created.ID {
			t.Errorf("Expected ID %d, got %d", created.ID, response.ID)
		}

		if response.Title != reqBody.Title {
			t.Errorf("Expected Title '%s', got '%s'", reqBody.Title, response.Title)
		}
	})

	t.Run("異常系: 存在しないIDの場合404", func(t *testing.T) {
		resp, err := httpClient.Get(fmt.Sprintf("%s/api/tags/999", testServerURL))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("異常系: 無効なIDの場合400", func(t *testing.T) {
		resp, err := httpClient.Get(fmt.Sprintf("%s/api/tags/invalid", testServerURL))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}

