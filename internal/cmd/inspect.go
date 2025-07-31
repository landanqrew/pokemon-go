package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/state"
)

func CommandInspect(cfg *config.Config) error {
	state := state.AppState
	pokemonName := cfg.Args[0]
	pokemon, err := state.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error getting pokemon: %v", err)
	}
	pokemon.PrintDescription()
	return nil
}