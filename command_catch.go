package main

import (
	"fmt"
	"math/rand"
)


func commandCatch(cfg *config, args ...string)error{
	if len(args) == 0 {
		return fmt.Errorf("A pokemon name is required after the 'catch' command")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %v... \n", pokemonName)

	res, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil{
		return err
	}

	catchProbability := rand.Intn(res.BaseExperience)
	threshold := 50

	if catchProbability < threshold{
		cfg.pokedex[pokemonName] = res	
		fmt.Printf("%v was caught! \n", pokemonName)
	}else{
		fmt.Printf("%v escaped! \n", pokemonName)
	}

	return nil
}