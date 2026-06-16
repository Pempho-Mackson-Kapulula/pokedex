package main

import (
	"fmt"
)


func commandMap(cfg *config, args ...string) error {
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


func commandMapBack (cfg *config,args ...string) error{
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
