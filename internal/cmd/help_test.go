package cmd

import (
	"io"
	"os"
	"testing"
)

func TestCommandHelp(t *testing.T) {
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
	}
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = originalStdout
	}()

	expectedStdout := `Welcome to the Pokedex!
Usage:

exit: Exit the Pokedex
help: Displays a help message

`

	err := CommandHelp()
	if err != nil {
		t.Errorf("CommandHelp() returned an error: %v", err)
	}

	w.Close()

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read from pipe: %v", err)
	}
	actualStdoutString := string(out)

	if actualStdoutString != expectedStdout {
		t.Errorf("CommandHelp() printed an unexpected output.\nGot:\n'%v'\nExpected:\n'%v'", []byte(actualStdoutString), []byte(expectedStdout))
	}
}
