package pokemon

import (
	"fmt"
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