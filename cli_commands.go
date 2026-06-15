package main

import (
	"fmt"
	"os"
	"github.com/Pempho-Mackson-Kapulula/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextUrl     string
	previousUrl string
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}, "map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    displayLocations,
		},"mapb":{
			name: "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon areas in the Pokemon world",
			callback: displayPreviousLocations,
		},
	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:`)

	for name, command := range getCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}
	return nil
}

func displayLocations(cfg *config) error {
	var pageURL *string
	if cfg.nextUrl != "" {
		pageURL = &cfg.nextUrl
	}

	response, err := cfg.pokeapiClient.ListLocations(pageURL)
	if err != nil {
		return err
	}

	
	if response.Next != nil {
		cfg.nextUrl = *response.Next
	} else {
		cfg.nextUrl = ""
	}

	if response.Previous != nil {
		cfg.previousUrl = *response.Previous
	} else {
		cfg.previousUrl = ""
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}


func displayPreviousLocations (cfg *config) error{
	var pageURL *string
	if cfg.previousUrl != "" {
		pageURL = &cfg.previousUrl
	}

	if cfg.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	response, err := cfg.pokeapiClient.ListLocations(pageURL)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	if response.Next != nil {
		cfg.nextUrl = *response.Next
	} else {
		cfg.nextUrl = ""
	}

	if response.Previous != nil {
		cfg.previousUrl = *response.Previous
	} else {
		cfg.previousUrl = ""
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}
