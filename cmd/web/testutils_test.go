package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EwokOwie/dog-api/internal/models"
)

// newTestApplication creates an application instance for testing.
func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	animals := models.NewAnimalService()

	return &application{
		logger:  logger,
		animals: animals,
	}
}

// testServer wraps httptest.Server with helper methods.
type testServer struct {
	*httptest.Server
}

// newTestServer creates a test server using the application routes.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	t.Helper()

	ts := httptest.NewServer(h)

	return &testServer{ts}
}

// get makes a GET request to the test server and returns the status code,
// headers, and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	t.Helper()

	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, string(body)
}
