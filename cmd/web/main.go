package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API v1 routes
	mux.HandleFunc("GET /api/v1/animals", handleListAnimals)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds", handleListBreeds)
	mux.HandleFunc("GET /api/v1/animals/{animal}/breeds/{breed}/photo", handleGetPhoto)

	// Health check
	mux.HandleFunc("GET /health", handleHealth)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
