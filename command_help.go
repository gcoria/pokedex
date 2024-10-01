package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("")
	fmt.Println("-----------------")
	fmt.Println("Avalaible Options")
	fmt.Println("-----------------")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s \n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Bye bye")
	os.Exit(0)
	return nil
}

func commandMap() error {
	return nil
}
