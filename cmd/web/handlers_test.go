package main

import (
	"net/http"
	"testing"

	"github.com/EwokOwie/dog-api/internal/assert"
)

func TestHealthCheck(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/health")

	assert.Equal(t, code, http.StatusOK)
	assert.StringContains(t, body, `"status":"ok"`)
}

func TestListAnimals(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/api/v1/animals")

	assert.Equal(t, code, http.StatusOK)
	assert.StringContains(t, body, `"data"`)
	assert.StringContains(t, body, `"dog"`)
}

func TestListBreeds(t *testing.T) {
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid animal",
			urlPath:  "/api/v1/animals/dog/breeds",
			wantCode: http.StatusOK,
			wantBody: `"data"`,
		},
		{
			name:     "Invalid animal",
			urlPath:  "/api/v1/animals/invalidanimal/breeds",
			wantCode: http.StatusNotFound,
			wantBody: `"error":"animal not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newTestApplication(t)
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.wantCode)
			assert.StringContains(t, body, tt.wantBody)
		})
	}
}

func TestGetPhoto(t *testing.T) {
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Invalid animal",
			urlPath:  "/api/v1/animals/invalidanimal/breeds/husky/photo",
			wantCode: http.StatusNotFound,
			wantBody: `"error":"animal not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newTestApplication(t)
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.wantCode)
			assert.StringContains(t, body, tt.wantBody)
		})
	}
}
