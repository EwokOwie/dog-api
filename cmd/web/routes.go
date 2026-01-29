package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// API v1 routes
	mux.HandleFunc("GET /api/v1/animals", app.handleListAnimals)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds", app.handleListBreeds)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds/{breed}/photo", app.handleGetPhoto)

	// Health check
	mux.HandleFunc("GET /health", app.handleHealth)

	return mux
}
