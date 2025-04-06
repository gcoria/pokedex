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
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid commands")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
			name:        "map_back",
			description: "get previous location",
			callback:    callbackMapPrevious,
		},
		"explore": {
			name:        "explore",
			description: "explore location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "inspect a pokemon of your pokemon collection",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show your pokemon collection",
			callback:    callbackPokedex,
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
