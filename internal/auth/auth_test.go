package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKeyValidKey(t *testing.T) {
	header := http.Header{"Authorization": []string{"ApiKey abc123"}}

	res, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if res != "abc1234" {
		t.Errorf("got %q, want abc123", res)
	}
}

func TestGetApiKeyNoHeader(t *testing.T) {
	header := http.Header{}

	res, err := GetAPIKey(header)

	if err == nil {
		t.Fatalf("expected error %v, got nil", ErrNoAuthHeaderIncluded)
	}
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("got err %v, want %v", err, ErrNoAuthHeaderIncluded)
	}

	if res != "" {
		t.Errorf("expected empty key on error, got %q", res)
	}
}
