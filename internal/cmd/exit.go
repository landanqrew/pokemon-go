package cmd

import (
	"fmt"
	"os"

	"github.com/landanqrew/pokemon-go/internal/config"
)

func CommandExit(cfg *config.Config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}