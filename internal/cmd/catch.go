package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/pokemon"
	"github.com/landanqrew/pokemon-go/internal/state"
)

func CommandCatch(cfg *config.Config) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.Args[0])
	state := state.AppState
	pokemonName := cfg.Args[0]
	url := cfg.Client.BaseURL + "pokemon/" + pokemonName
	response, err := cfg.Client.Cache.Get(url)
	if err != nil {
		return fmt.Errorf("error getting response at url (%s): %v", url, err)
	}
	pokemon, err := pokemon.ParsePokemonFromResponse(response)
	if err != nil {
		return fmt.Errorf("error parsing pokemon: %v", err)
	}

	if pokemon.Catch() {
		fmt.Printf("%s was caught!\n", pokemonName)
		state.AddPokemon(pokemonName, pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}