package pokemon

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetAndStoreLocationAreas(t *testing.T) {
	locationAreas, err := GetAndStoreLocationAreas()
	if err != nil {
		t.Errorf("Error getting location areas: %v", err)
	}

	if len(locationAreas) == 0 {
		t.Errorf("No location areas found")
	} else if len(locationAreas) != 1089 {
		t.Errorf("Expected 1089 location areas, got %d", len(locationAreas))
	}
}

func TestGetLocationNames(t *testing.T) {
	locationNames, err := ReadLocationsFromCache()
	if err != nil {
		t.Errorf("Error getting location names: %v", err)
	}
	fmt.Printf("locationNames: %d\n", len(locationNames))
}

func TestGetPokemonNames(t *testing.T) {
	testLocations := []struct{
		URL string
		Name string
		ExpectsError bool
		ExpectedPokemonNames []string
	}{
		{
			URL: "https://pokeapi.co/api/v2/location-area/",
			Name: "pastoria-city-area",
			ExpectsError: false,
			ExpectedPokemonNames: []string{"tentacool","tentacruel","magikarp"},
		},
	}

	locationAreas := []LocationArea{}
	for _, test := range testLocations {
		locationAreas = append(locationAreas, LocationArea{
			URL: test.URL,
			Name: test.Name,
		})
	}

	for i, locationArea := range locationAreas {
		pokemonNames, err := locationArea.GetPokemonNames()
		if err != nil && !testLocations[i].ExpectsError {
			t.Errorf("expected no error, got %v", err)
		}
		if err == nil && testLocations[i].ExpectsError {
			t.Errorf("expected error, got none")
		}
		for _, p := range testLocations[i].ExpectedPokemonNames {
			if !slices.Contains(pokemonNames, p) {
				t.Errorf("expected %v to be in %v", p, pokemonNames)
			}
		}
	}

	
}