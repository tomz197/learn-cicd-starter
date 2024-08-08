package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("no auth header", func(t *testing.T) {
		headers := make(http.Header)
		apiKey, err := GetAPIKey(headers)
		if apiKey != "" {
			t.Errorf("expected empty string, got %#v", apiKey)
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected ErrNoAuthHeaderIncluded, got %#v", err)
		}
	})

	t.Run("malformed auth header", func(t *testing.T) {
		headers := make(http.Header)
		headers.Set("Authorization", "Bearer")
		apiKey, err := GetAPIKey(headers)
		if apiKey != "" {
			t.Errorf("expected empty string, got %#v", apiKey)
		}
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("valid auth header", func(t *testing.T) {
		headers := make(http.Header)
		headers.Set("Authorization", "ApiKey abc123")
		apiKey, err := GetAPIKey(headers)
		if apiKey != "abc123" {
			t.Errorf("expected abc123, got %#v", apiKey)
		}
		if err != nil {
			t.Errorf("expected nil, got %#v", err)
		}
	})
}
