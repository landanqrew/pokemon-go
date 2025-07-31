package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/state"
)

func CommandList(cfg *config.Config) error {
	state := state.AppState
	pokemonNames := state.ListPokemon()
	if len(pokemonNames) == 0 {
		fmt.Println("No pokemon in the pokedex")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemonName := range pokemonNames {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}