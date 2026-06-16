package main

import (
	"fmt"
)


func commandInspect(cfg *config, args ...string)error{
	if len(args) == 0 {
		return fmt.Errorf("A pokemon name is required after the 'catch' command")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]

	if !ok{
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats{
		fmt.Printf("   - %v : %v\n" , stat.Stat.Name , stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, typeInfo := range pokemon.Types{
		fmt.Printf("   - %v\n", typeInfo.Type.Name)
	}

	
	return nil
}