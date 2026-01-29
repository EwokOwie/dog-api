package models

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EwokOwie/dog-api/internal/assert"
)

// mockDogAPI creates a mock server that simulates the Dog CEO API.
func mockDogAPI(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()

	// Mock breeds list endpoint
	mux.HandleFunc("/api/breeds/list/all", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": map[string][]string{
				"husky":    {},
				"labrador": {"yellow", "black"},
				"poodle":   {"toy", "standard"},
			},
			"status": "success",
		})
	})

	// Mock breed image endpoint
	mux.HandleFunc("/api/breed/husky/images/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "https://images.dog.ceo/breeds/husky/test.jpg",
			"status":  "success",
		})
	})

	// Mock invalid breed (404)
	mux.HandleFunc("/api/breed/invalidbreed/images/random", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Breed not found",
			"status":  "error",
		})
	})

	return httptest.NewServer(mux)
}

// newTestDog creates a Dog instance pointing to the mock server.
func newTestDog(t *testing.T, baseURL string) *Dog {
	t.Helper()

	dog := NewDog()
	// We need to override the base URL, but it's a const.
	// For testing, we'll create a custom dog with the test URL.
	return &Dog{
		client:  dog.client,
		baseURL: baseURL + "/api",
	}
}

func TestDogName(t *testing.T) {
	dog := NewDog()
	assert.Equal(t, dog.Name(), "dog")
}

func TestDogGetBreeds(t *testing.T) {
	server := mockDogAPI(t)
	defer server.Close()

	dog := newTestDog(t, server.URL)
	breeds, err := dog.GetBreeds()

	assert.NilError(t, err)

	// Check that we got the expected breeds
	if len(breeds) != 3 {
		t.Errorf("got %d breeds; want 3", len(breeds))
	}

	// Breeds should be sorted
	assert.Equal(t, breeds[0], "husky")
	assert.Equal(t, breeds[1], "labrador")
	assert.Equal(t, breeds[2], "poodle")
}

func TestDogGetBreedPhoto(t *testing.T) {
	tests := []struct {
		name      string
		breed     string
		wantURL   string
		wantError error
	}{
		{
			name:      "Valid breed",
			breed:     "husky",
			wantURL:   "https://images.dog.ceo/breeds/husky/test.jpg",
			wantError: nil,
		},
		{
			name:      "Invalid breed",
			breed:     "invalidbreed",
			wantURL:   "",
			wantError: ErrBreedNotFound,
		},
	}

	server := mockDogAPI(t)
	defer server.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dog := newTestDog(t, server.URL)
			url, err := dog.GetBreedPhoto(tt.breed)

			if tt.wantError != nil {
				if err == nil {
					t.Errorf("expected error %v; got nil", tt.wantError)
				}
			} else {
				assert.NilError(t, err)
				assert.Equal(t, url, tt.wantURL)
			}
		})
	}
}
