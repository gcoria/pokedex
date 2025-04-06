# Pokédex CLI

A feature-rich command-line interface (CLI) application written in Go that simulates a Pokédex from the Pokémon universe. This application interacts with the [PokéAPI](https://pokeapi.co/) to provide various Pokémon-related functionalities.

![Pokédex Demo](https://via.placeholder.com/800x400?text=Pok%C3%A9dex+CLI+Demo)

## Features

- **Explore Location Areas**: Navigate through the Pokémon world by exploring different location areas
- **Catch Pokémon**: Attempt to catch Pokémon with a success rate based on their base experience
- **Pokédex Collection**: View and manage your collection of caught Pokémon
- **Pokémon Details**: Inspect detailed information about your caught Pokémon, including:
  - Base stats
  - Abilities
  - Types
  - Moves
  - Physical attributes (height, weight)
- **Smart Caching**: Implements efficient caching to minimize API calls and improve performance

## Technologies Used

- **Go (Golang)**: Core programming language
- **PokéAPI**: External API for Pokémon data
- **Custom Cache Implementation**: Time-based cache with automatic expiration
- **Concurrent Programming**: Uses goroutines and mutexes for thread-safe caching
- **Unit Testing**: Test coverage for critical components

## Installation

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository

   ```bash
   git clone https://github.com/yourusername/pokedex.git
   cd pokedex
   ```

2. Build the application

   ```bash
   go build -o pokedex
   ```

3. Run the application
   ```bash
   ./pokedex
   ```

## Usage

After launching the application, you'll be presented with a command prompt `pokedex>`. Here are the available commands:

- `help`: Display a list of available commands
- `map`: View the next set of location areas
- `map_back`: View the previous set of location areas
- `explore`: Explore Pokémon in the current location area
- `catch <pokemon>`: Attempt to catch a specific Pokémon
- `inspect <pokemon>`: View detailed information about a caught Pokémon
- `pokedex`: View your collection of caught Pokémon
- `exit`: Exit the application

### Example Usage

```bash
pokedex> map
Location Areas:
 - canalave-city-area
 - eterna-city-area
 - pastoria-city-area
 - sunyshore-city-area

pokedex> explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacruel
 - staryu
 - magikarp

pokedex> catch staryu
You caught the pokemon: staryu

pokedex> inspect staryu
Name: staryu
Height: 8
Weight: 345
Base Experience: 68
Abilities:
 - illuminate
 - natural-cure
Stats:
 - hp: 30
 - attack: 45
 - defense: 55
Types:
 - water
```

## Architecture

The application follows a modular design with clear separation of concerns:

- **Main Package**: Core application logic and REPL (Read-Eval-Print Loop) implementation
- **Command Handlers**: Individual command implementations
- **PokéAPI Client**: Handles communication with the external API
- **Cache Package**: Implements efficient caching with automatic expiration

### Cache Implementation

The project features a custom cache implementation with the following characteristics:

- Thread-safe with mutex locks
- Automatic expiration of cached items
- Background goroutine for cache cleanup

## Project Goals

This project was developed with the following objectives:

- Learn JSON parsing in Go
- Practice making HTTP requests in Go
- Build a CLI tool that makes interacting with a back-end server easier
- Get hands-on practice with local Go development and tooling
- Learn about caching and how to use it to improve performance

## Future Enhancements

- Add persistent storage to save caught Pokémon between sessions
- Implement battles between caught Pokémon
- Add more interactive commands and features
- Expand the UI with color and ASCII art
- Allow Evolutions

## Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the Pokémon data API
- The Go community for excellent documentation and resources
