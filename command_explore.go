package main

import (
	"fmt"
)

func commandExplore (cfg *config, args ...string)error{
	areaName := args[0]


	location , err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil{
		return err
	}

	fmt.Printf("Exploring %v...",areaName)
	fmt.Println("Found Pokemon:")

	
	for _, encounter := range location.PokemonEncounters{
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}