package cmd

type CliCommand struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Callback    func() error `json:"-"`
}

var CommandListMap = map[string]CliCommand{}

func BuildCommandMap() {
	CommandListMap = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays a map of the Pokedex",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous page of the map",
			Callback:    CommandMapBack,
		},
	}
}
