

# Pokedex CLI

A lightweight, terminal-based Pokedex application built in Go. This project features a custom REPL (Read-Eval-Print Loop) to interact with the [PokeAPI](https://pokeapi.co/).

## Features

- **Real-time Data**: Fetch location areas and Pokémon stats directly from the PokeAPI.
- **Caching**: Optimized performance using a custom-built, thread-safe cache system to reduce network requests.
- **Advanced REPL**: 
    - Interactive command-line interface.
    - Persistent command history (even after closing the app!).
    - Smooth navigation using Up/Down arrow keys.
- **Pokedex System**: Catch Pokémon and store them in your personal Pokedex to inspect their stats later.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/pokedex

2. Build the project:

    go build -o pokedex

3. Run the application:

    ./pokedex


