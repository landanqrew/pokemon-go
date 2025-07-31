package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/pokemon"
	"github.com/landanqrew/pokemon-go/internal/state"
)


func CommandMap(cfg *config.Config) error {
	// fmt.Println("Map")
	locationNames, err := pokemon.GetLocationNames()
	if err != nil {
		return fmt.Errorf("error getting location names: %v", err)
	}

	appState := state.AppState
	// fmt.Printf("appState.Page: %d\n", appState.LocationNamePage)
	if appState.LocationNamePage == 0 || appState.LocationNamePage*20 > len(locationNames) {
		appState.ResetLocationPage()
	} else {
		appState.IncrementLocationPage()
		// fmt.Printf("appState.Page: %d\n", appState.LocationNamePage)
	}
	page := appState.GetLocationNamePage()
	// fmt.Printf("page: %d\n", page)

	startIndex := (page - 1) * 20
	endIndex := startIndex + 20

	if endIndex > len(locationNames) {
		endIndex = len(locationNames)
	}

	for _, locationName := range locationNames[startIndex:endIndex] {
		fmt.Println(locationName)
	}

	return nil
}
