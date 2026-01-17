package auth_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	authRequest "github.com/jobpay/todo/internal/presentation/request/auth"
	authResponse "github.com/jobpay/todo/internal/presentation/response/auth"
)

func TestGetMe(t *testing.T) {
	cleanupDB()

	t.Run("正常系: ユーザー情報を取得できる", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "getme@example.com",
			Name:  "GetMe User",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, _ := httpClient.Post(
			testServerURL+"/api/auth/register",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		var registerResponse authResponse.UserResponse
		json.NewDecoder(resp.Body).Decode(&registerResponse)
		resp.Body.Close()

		resp2, err := httpClient.Get(
			fmt.Sprintf("%s/api/auth/me?user_id=%d", testServerURL, registerResponse.ID),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp2.Body.Close()

		if resp2.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp2.StatusCode)
		}

		var response authResponse.UserResponse
		if err := json.NewDecoder(resp2.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Email != reqBody.Email {
			t.Errorf("Expected Email '%s', got '%s'", reqBody.Email, response.Email)
		}

		if response.Name != reqBody.Name {
			t.Errorf("Expected Name '%s', got '%s'", reqBody.Name, response.Name)
		}
	})

	t.Run("異常系: user_idが指定されていない場合エラー", func(t *testing.T) {
		resp, err := httpClient.Get(
			testServerURL + "/api/auth/me",
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("異常系: 存在しないユーザーID", func(t *testing.T) {
		resp, err := httpClient.Get(
			fmt.Sprintf("%s/api/auth/me?user_id=99999", testServerURL),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
