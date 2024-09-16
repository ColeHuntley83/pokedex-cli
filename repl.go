package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringPrompt(label string) (string, error) {

	var output string
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, label+" ")
		output, _ = input.ReadString('\n')

		if output != "" {
			break
		}
	}
	return strings.TrimSpace(output), nil

}

func runRepl(config *config) {
	fmt.Println("waiting for input")
	for {

		userInput, _ := stringPrompt("Pokedex >")
		if userInput == "" {
			continue
		}
		commands := GetPromptCommands()
		value, ok := commands[userInput]
		if !ok {
			fmt.Println("command not found options [help | exit]")
			continue
		}
		err := value.callBack(config)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}

}
