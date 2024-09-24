package main

import "fmt"


func helpMenu() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "displays help message",
			callback: commandHelp
		},
		"exit": {
			name: "exit",
			description: "exit cli menu",
			callback: commandExit
		}
	}
}