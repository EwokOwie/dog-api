package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleListAnimals(w http.ResponseWriter, r *http.Request) {
	log.Println("handleListAnimals called")
	fmt.Fprintln(w, "hello from handleListAnimals")
}

func handleListBreeds(w http.ResponseWriter, r *http.Request) {
	animal := r.PathValue("animal")
	log.Printf("handleListBreeds called for animal: %s", animal)
	fmt.Fprintf(w, "hello from handleListBreeds - animal: %s\n", animal)
}

func handleGetPhoto(w http.ResponseWriter, r *http.Request) {
	animal := r.PathValue("animal")
	breed := r.PathValue("breed")
	log.Printf("handleGetPhoto called for animal: %s, breed: %s", animal, breed)
	fmt.Fprintf(w, "hello from handleGetPhoto - animal: %s, breed: %s\n", animal, breed)
}

func handleGetSubBreedPhoto(w http.ResponseWriter, r *http.Request) {
	animal := r.PathValue("animal")
	breed := r.PathValue("breed")
	subbreed := r.PathValue("subbreed")
	log.Printf("handleGetSubBreedPhoto called for animal: %s, breed: %s, subbreed: %s", animal, breed, subbreed)
	fmt.Fprintf(w, "hello from handleGetSubBreedPhoto - animal: %s, breed: %s, subbreed: %s\n", animal, breed, subbreed)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("handleHealth called")
	fmt.Fprintln(w, "hello from handleHealth")
}
