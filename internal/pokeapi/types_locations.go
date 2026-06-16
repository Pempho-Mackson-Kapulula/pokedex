package pokeapi

type Response struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type PokemonEncounter struct {
    Pokemon Pokemon `json:"pokemon"`
}

type Location struct {
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
