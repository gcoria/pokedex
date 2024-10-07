package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid commands")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "get location near",
			callback:    callbackMap,
		},
		"map_back": {
			name:        "map",
			description: "get previous location",
			callback:    callbackMapPrevious,
		},
		"exit": {
			name:        "exit",
			description: "exit cli menu",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
