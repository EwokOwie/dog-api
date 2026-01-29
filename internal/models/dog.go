package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

const dogAPIBaseURL = "https://dog.ceo/api"

// dogAPIResponse represents the standard response format from dog.ceo API
type dogAPIResponse struct {
	Message json.RawMessage `json:"message"`
	Status  string          `json:"status"`
}

// Dog implements the Animal interface for dog breeds
type Dog struct {
	client *http.Client
}

// NewDog creates a new Dog instance with a configured HTTP client
func NewDog() *Dog {
	return &Dog{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Name returns the animal type name
func (d *Dog) Name() string {
	return "dog"
}

// GetBreeds fetches all available dog breeds from the Dog API
func (d *Dog) GetBreeds() ([]string, error) {
	resp, err := d.client.Get(dogAPIBaseURL + "/breeds/list/all")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch breeds: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dog API returned status %d", resp.StatusCode)
	}

	var apiResp dogAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Status != "success" {
		return nil, fmt.Errorf("dog API returned status: %s", apiResp.Status)
	}

	// The message is a map of breed -> sub-breeds
	var breedsMap map[string][]string
	if err := json.Unmarshal(apiResp.Message, &breedsMap); err != nil {
		return nil, fmt.Errorf("failed to parse breeds: %w", err)
	}

	breeds := make([]string, 0, len(breedsMap))
	for breed := range breedsMap {
		breeds = append(breeds, breed)
	}
	sort.Strings(breeds)

	return breeds, nil
}

// GetBreedPhoto fetches a random photo URL for the specified breed
func (d *Dog) GetBreedPhoto(breed string) (string, error) {
	url := fmt.Sprintf("%s/breed/%s/images/random", dogAPIBaseURL, breed)
	resp, err := d.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch photo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("dog API returned status %d", resp.StatusCode)
	}

	var apiResp dogAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Status != "success" {
		return "", fmt.Errorf("dog API returned status: %s", apiResp.Status)
	}

	var photoURL string
	if err := json.Unmarshal(apiResp.Message, &photoURL); err != nil {
		return "", fmt.Errorf("failed to parse photo URL: %w", err)
	}

	return photoURL, nil
}
