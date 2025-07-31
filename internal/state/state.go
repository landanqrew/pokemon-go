package state

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/pokemon"
)

// Global state instance
var AppState = &State{}

type State struct {
	LocationNamePage int `json:"locationNamePage"`
	Pokedex map[string]pokemon.Pokemon
}

func (s *State) GetLocationNamePage() int {
	if s.LocationNamePage == 0 {
		s.LocationNamePage = 1
	}
	return s.LocationNamePage
}

func (s *State) IncrementLocationPage() {
	s.LocationNamePage += 1
}

func (s *State) DecrementLocationPage() {
	s.LocationNamePage -= 1
}

func (s *State) ResetLocationPage() {
	s.LocationNamePage = 1
}

func (s *State) Init() {
	s.LocationNamePage = 0
	s.Pokedex = make(map[string]pokemon.Pokemon)
}

func (s *State) AddPokemon(pokemonName string, pokemon pokemon.Pokemon) {
	s.Pokedex[pokemonName] = pokemon
}

func (s *State) GetPokemon(pokemonName string) (pokemon.Pokemon, error) {
	nilPokemon := pokemon.Pokemon{}
	pokemon, ok := s.Pokedex[pokemonName]
	if !ok {
		return nilPokemon, fmt.Errorf("you have not caught that pokemon")
	}
	return pokemon, nil
}

func (s *State) ListPokemon() []string {
	pokemonNames := []string{}
	for pokemonName := range s.Pokedex {
		pokemonNames = append(pokemonNames, pokemonName)
	}
	return pokemonNames
}
