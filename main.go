package main

import (
	"time"

	"github.com/ericsprague3/pokedexcli/internal/pokecache"
)

func main() {
	cfg := Config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Provide help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "List Location areas of the map",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "List the last 20 Location areas of the map",
				callback:    commandMapBack,
			},
			"explore": {
				name:        "explore",
				description: "Explore location area",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Attempt to catch a Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a Pokemon",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Inspect caught Pokemon",
				callback:    commandPokedex,
			},
		},
		cache:   pokecache.NewCache(time.Minute * 5),
		pokemon: make(map[string]PokemonResult),
	}

	startRepl(&cfg)
}
