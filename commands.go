package main

import (
	"fmt"
	"os"
)



type command struct {
 name string
 description string
 callBack func(*config) error
 
}


func  helpCommand(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, command := range GetPromptCommands(){
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	fmt.Println()
	
	return nil
}


func exitCommand(config *config) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)
	return nil
}

func GetPromptCommands() map[string]command {

 commands := map[string]command{
	"help" : {
		name: "help",
		description: "Displays a help message",
		callBack: helpCommand,
	},
	"exit" : {
		name: "exit",
		description: "quits the pokedex",
		callBack: exitCommand,
	},
	"map" : {
		name: "map",
		description: "displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations",
		callBack: mapCommand,
	},
 }
 return commands

}