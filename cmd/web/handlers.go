package main

import (
	"errors"
	"net/http"

	"github.com/EwokOwie/dog-api/internal/models"
)

func (app *application) handleListAnimals(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("listing animals")
	animals := app.animals.ListAnimals()
	app.writeJSON(w, http.StatusOK, animals)
}

func (app *application) handleListBreeds(w http.ResponseWriter, r *http.Request) {
	animalName := r.PathValue("animal")
	app.logger.Info("listing breeds", "animal", animalName)

	animal, ok := app.animals.Get(animalName)
	if !ok {
		app.notFound(w, "animal not found")
		return
	}

	breeds, err := animal.GetBreeds()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, breeds)
}

func (app *application) handleGetPhoto(w http.ResponseWriter, r *http.Request) {
	animalName := r.PathValue("animal")
	breed := r.PathValue("breed")
	app.logger.Info("getting photo", "animal", animalName, "breed", breed)

	animal, ok := app.animals.Get(animalName)
	if !ok {
		app.notFound(w, "animal not found")
		return
	}

	photoURL, err := animal.GetBreedPhoto(breed)
	if err != nil {
		if errors.Is(err, models.ErrBreedNotFound) {
			app.notFound(w, "breed not found")
			return
		}
		app.serverError(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, map[string]string{"url": photoURL})
}

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("health check")
	app.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
