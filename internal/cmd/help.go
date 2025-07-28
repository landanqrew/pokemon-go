package cmd

import "fmt"


func CommandHelp() error {
	BuildCommandMap()
	printStr := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, v := range CommandListMap {
		printStr += fmt.Sprintf("%s: %s\n", v.Name, v.Description)
	}
	fmt.Println(printStr)
	return nil
}