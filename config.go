package main

import (
	"github.com/ericsprague3/pokedexcli/internal/pokecache"
)

type Config struct {
	commands map[string]cliCommand
	next     string
	previous string
	cache    *pokecache.Cache
	pokemon  map[string]PokemonResult
}
