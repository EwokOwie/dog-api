package models

// Animal defines the interface that all animal types must implement.
// This allows for polymorphic behavior when adding new animals.
type Animal interface {
	// Name returns the animal type name (e.g., "dog", "cat")
	Name() string
	// GetBreeds returns all available breeds for this animal
	GetBreeds() ([]string, error)
	// GetBreedPhoto returns a photo URL for the specified breed
	GetBreedPhoto(breed string) (string, error)
}

// AnimalService manages registered animal providers
type AnimalService struct {
	registry map[string]Animal
}

// NewAnimalService creates a new AnimalService with default providers registered
func NewAnimalService() *AnimalService {
	s := &AnimalService{
		registry: make(map[string]Animal),
	}

	// Register default providers
	s.Register(NewDog())

	return s
}

// Register adds an animal provider to the service
func (s *AnimalService) Register(a Animal) {
	s.registry[a.Name()] = a
}

// Get retrieves an animal by name
func (s *AnimalService) Get(name string) (Animal, bool) {
	a, ok := s.registry[name]
	return a, ok
}

// ListAnimals returns all registered animal names
func (s *AnimalService) ListAnimals() []string {
	names := make([]string, 0, len(s.registry))
	for name := range s.registry {
		names = append(names, name)
	}
	return names
}
