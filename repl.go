package main

import (
	"strings"
	"github.com/Pempho-Mackson-Kapulula/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
	
}

type config struct {
	pokeapiClient pokeapi.Client
	nextUrl       string
	previousUrl   string
	pokedex 	  map[string]pokeapi.PokemonData
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
			callback:    commandMap,
		},"mapb":{
			name: "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon areas in the Pokemon world",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Displays all pokemons in a given situation",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a Pokemon and store it in your Pokedex",
			callback: commandCatch,
		},
	}
}


func cleanInput(text string) []string{
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}