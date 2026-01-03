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

func TestDeleteTag(t *testing.T) {
	cleanupDB()

	t.Run("正常系: タグを削除できる", func(t *testing.T) {
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

		// 削除
		req, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/api/tags/%d", testServerURL, created.ID),
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

		// DBから削除されているか確認
		var count int64
		testDB.Table("tags").Where("id = ?", created.ID).Count(&count)
		if count != 0 {
			t.Errorf("Expected 0 records in DB, got %d", count)
		}
	})

	t.Run("異常系: 存在しないIDの場合404", func(t *testing.T) {
		req, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/api/tags/999", testServerURL),
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

	t.Run("異常系: 無効なIDの場合400", func(t *testing.T) {
		req, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/api/tags/invalid", testServerURL),
			nil,
		)

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

