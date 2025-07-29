package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/pokemon"
	"github.com/landanqrew/pokemon-go/internal/state"
)


func CommandMapBack() error {
	// fmt.Println("Mapb")
	locationNames, err := pokemon.GetLocationNames()
	if err != nil {
		return fmt.Errorf("error getting location names: %v", err)
	}

	appState := state.AppState
	if appState.GetLocationNamePage() <= 1 {
		fmt.Println("cannot decrement page. Input 'map' to go to the next page")
		appState.ResetLocationPage()
	} else {
		appState.DecrementLocationPage()
	}
	page := appState.GetLocationNamePage()

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