package cmd

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/config"
)


func CommandHelp(cfg *config.Config) error {
	BuildCommandMap()
	printStr := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, v := range CommandListMap {
		printStr += fmt.Sprintf("%s: %s\n", v.Name, v.Description)
	}
	fmt.Println(printStr)
	return nil
}