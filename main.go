package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/landanqrew/pokemon-go/internal/cmd"
	"github.com/landanqrew/pokemon-go/internal/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd.BuildCommandMap() // initialize command map

	// inputs := []string{}

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := utils.CleanInput(text)
		if len(cleaned) > 0 {
			command := strings.ToLower(cleaned[0])
			if _, ok := cmd.CommandListMap[command]; ok {
				cmd.CommandListMap[command].Callback()
			} else {
				fmt.Printf("Your command was: %v. This is not a valid command\n", command)
				cmd.CommandHelp()
			}
		}
	}
}