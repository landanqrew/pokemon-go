package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/pokemon"
)


func CommandExplore(cfg *config.Config) error {
	location := cfg.Args[0]
	locationArea := pokemon.LocationArea{
		Name: location,
		URL: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location),
	}

	pokemonNames, err := locationArea.GetPokemonNames()
	if err != nil {
		return fmt.Errorf("error getting pokemon names: %v", err)
	}

	fmt.Printf("Exploring %s...\n", location)
	for _, pokemonName := range pokemonNames {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}