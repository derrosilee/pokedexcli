package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		switch command {
		case "help":
			fmt.Println("Welcome to Pokedex help menu!")
			fmt.Println("Here are your available commands:")
			fmt.Println(" - help")
			fmt.Println(" - exit")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
		}
	}

}

type cliCommands struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns of the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
