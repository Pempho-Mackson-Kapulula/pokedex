package pokeapi

type PokemonData struct {
    Name           string        `json:"name"`
    BaseExperience int           `json:"base_experience"`
    Height         int           `json:"height"`
    Weight         int           `json:"weight"`
    Stats          []PokemonStat `json:"stats"`
    Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}


type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

