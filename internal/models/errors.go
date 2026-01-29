package models

import "errors"

var (
	// ErrAnimalNotFound is returned when the requested animal type is not registered
	ErrAnimalNotFound = errors.New("animal not found")

	// ErrBreedNotFound is returned when the requested breed does not exist
	ErrBreedNotFound = errors.New("breed not found")

	// ErrUpstreamAPI is returned when the external API fails
	ErrUpstreamAPI = errors.New("upstream API error")
)
