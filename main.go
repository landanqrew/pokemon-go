package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/landanqrew/pokemon-go/internal/api"
	"github.com/landanqrew/pokemon-go/internal/cmd"
	"github.com/landanqrew/pokemon-go/internal/config"
	"github.com/landanqrew/pokemon-go/internal/pokecache"
	"github.com/landanqrew/pokemon-go/internal/utils"
)

var cachePtr *pokecache.Cache = pokecache.NewCache(10 * time.Minute)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd.BuildCommandMap() // initialize command map
	config := config.GetConfig()
	config.Client = api.NewClient(cachePtr, "https://pokeapi.co/api/v2/")



	// inputs := []string{}

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := utils.CleanInput(text)
		if len(cleaned) > 0 {
			command := strings.ToLower(cleaned[0])
			if _, ok := cmd.CommandListMap[command]; ok {
				if len(cmd.CommandListMap[command].Args) > 0 {
					if len(cleaned) == 1 {
						config.Args = []string{}
						fmt.Printf("You need to provide an argument after %s for the %s command\n", command, command)
					} else {
						config.Args = cleaned[1:]
						cmd.CommandListMap[command].Callback(config)
					}
				} else {
					cmd.CommandListMap[command].Callback(config)
				}
				
			} else {
				fmt.Printf("Your command was: %v. This is not a valid command\n", command)
				cmd.CommandHelp(config)
			}
		}
	}
}