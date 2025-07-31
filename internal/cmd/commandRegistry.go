package cmd

import "github.com/landanqrew/pokemon-go/internal/config"

type CliCommand struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Args        []string     `json:"args"`
	Callback    func(*config.Config) error `json:"-"`
}

var CommandListMap = map[string]CliCommand{}

func BuildCommandMap() {
	CommandListMap = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Args:        []string{},
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Args:        []string{},
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays a map of the Pokedex",
			Args:        []string{},
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous page of the map",
			Args:        []string{},
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays a list of pokemon in a location",
			Args:        []string{"location"},
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Percentage chance of catching a pokemon and adding it to the pokedex",
			Args:        []string{"pokemonName"},
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays a list of pokemon in the pokedex",
			Args:        []string{},
			Callback:    CommandList,
		},
		"list": {
			Name:        "list",
			Description: "Displays a list of pokemon in the pokedex",
			Args:        []string{},
			Callback:    CommandList,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Displays a detailed description of a pokemon",
			Args:        []string{"pokemonName"},
			Callback:    CommandInspect,
		},
	}
}
