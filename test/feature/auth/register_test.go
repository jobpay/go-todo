package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	authRequest "github.com/jobpay/todo/internal/presentation/request/auth"
	authResponse "github.com/jobpay/todo/internal/presentation/response/auth"
)

func TestRegister(t *testing.T) {
	cleanupDB()

	t.Run("正常系: ユーザーを登録できる", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "test@example.com",
			Name:  "Test User",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/auth/register",
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

		var response authResponse.UserResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.ID <= 0 {
			t.Errorf("Expected ID > 0, got %d", response.ID)
		}

		if response.Email != reqBody.Email {
			t.Errorf("Expected Email '%s', got '%s'", reqBody.Email, response.Email)
		}

		if response.Name != reqBody.Name {
			t.Errorf("Expected Name '%s', got '%s'", reqBody.Name, response.Name)
		}

		var count int64
		testDB.Table("users").Where("id = ?", response.ID).Count(&count)
		if count != 1 {
			t.Errorf("Expected 1 record in DB, got %d", count)
		}
	})

	t.Run("異常系: メールアドレスが空の場合エラー", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "",
			Name:  "Test User",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/auth/register",
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

	t.Run("異常系: メールアドレスが不正な形式", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "invalid-email",
			Name:  "Test User",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/auth/register",
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

	t.Run("異常系: 名前が空の場合エラー", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "test@example.com",
			Name:  "",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp, err := httpClient.Post(
			testServerURL+"/api/auth/register",
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

	t.Run("異常系: 同じメールアドレスで登録できない", func(t *testing.T) {
		reqBody := authRequest.RegisterRequest{
			Email: "duplicate@example.com",
			Name:  "User 1",
		}
		jsonBody, _ := json.Marshal(reqBody)

		resp1, _ := httpClient.Post(
			testServerURL+"/api/auth/register",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		resp1.Body.Close()

		resp2, err := httpClient.Post(
			testServerURL+"/api/auth/register",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp2.Body.Close()

		if resp2.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp2.StatusCode)
		}
	})
}
