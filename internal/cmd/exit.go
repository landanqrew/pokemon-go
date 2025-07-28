package cmd

import (
	"fmt"
	"os"
)

func CommandExit() error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}