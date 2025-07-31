package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/pokemon"
)


func CommandExplore(cfg *config.Config) error {
	location := cfg.Args[0]
	response, err := cfg.Client.Cache.Get(cfg.Client.BaseURL + "location-area/" + location)
	if err != nil {
		return fmt.Errorf("error getting location area details: %v", err)
	}

	pokemonNames, err := pokemon.ParsePokemonNamesFromResponse(response)
	if err != nil {
		return fmt.Errorf("error parsing pokemon names: %v", err)
	}

	fmt.Printf("Exploring %s...\n", location)
	for _, pokemonName := range pokemonNames {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}