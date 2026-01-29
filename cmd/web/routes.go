package main

import (
	"net/http"

	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// API v1 routes
	mux.HandleFunc("GET /api/v1/animals", app.handleListAnimals)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds", app.handleListBreeds)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds/{breed}/photo", app.handleGetPhoto)

	// Health check
	mux.HandleFunc("GET /health", app.handleHealth)

	// Serve OpenAPI spec
	mux.HandleFunc("GET /api/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/openapi.yaml")
	})

	// Swagger UI
	mux.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/api/openapi.yaml"),
	))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
