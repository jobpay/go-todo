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

func TestUpdateTag(t *testing.T) {
	cleanupDB()

	t.Run("正常系: タグを更新できる", func(t *testing.T) {
		// テストデータを作成
		createReq := tagRequest.StoreRequest{
			Title: "urgent",
		}
		jsonBody, _ := json.Marshal(createReq)

		createResp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		defer createResp.Body.Close()

		var created tagResponse.TagResponse
		json.NewDecoder(createResp.Body).Decode(&created)

		// 更新
		updateReq := tagRequest.UpdateRequest{
			Title: "high-priority",
		}
		jsonBody, _ = json.Marshal(updateReq)

		req, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("%s/api/tags/%d", testServerURL, created.ID),
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

		var response tagResponse.TagResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Title != updateReq.Title {
			t.Errorf("Expected Title '%s', got '%s'", updateReq.Title, response.Title)
		}
	})

	t.Run("異常系: 存在しないIDの場合404", func(t *testing.T) {
		updateReq := tagRequest.UpdateRequest{
			Title: "test",
		}
		jsonBody, _ := json.Marshal(updateReq)

		req, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("%s/api/tags/999", testServerURL),
			bytes.NewBuffer(jsonBody),
		)
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("異常系: タイトルが空の場合エラー", func(t *testing.T) {
		// テストデータを作成
		createReq := tagRequest.StoreRequest{
			Title: "urgent",
		}
		jsonBody, _ := json.Marshal(createReq)

		createResp, _ := httpClient.Post(
			testServerURL+"/api/tags",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		defer createResp.Body.Close()

		var created tagResponse.TagResponse
		json.NewDecoder(createResp.Body).Decode(&created)

		// 更新
		updateReq := tagRequest.UpdateRequest{
			Title: "",
		}
		jsonBody, _ = json.Marshal(updateReq)

		req, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("%s/api/tags/%d", testServerURL, created.ID),
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

