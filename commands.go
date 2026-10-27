package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(c *Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(c *Config, _ string) error {
	data := getLocationAreas(c, false)
	c.next = data.Next
	c.previous = data.Previous
	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapBack(c *Config, _ string) error {
	data := getLocationAreas(c, true)
	c.next = data.Next
	c.previous = data.Previous
	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandExplore(c *Config, name string) error {
	fmt.Printf("Exploring %s\n", name)
	data := getLocationAreaByName(c, name)
	pokemonEncounters := data.PokemonEncounters
	for _, v := range pokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *Config, name string) error {
	if name == "" {
		return errors.New("Pokemon name not provided")
	}

	_, alreadyCaught := c.pokemon[name]
	if alreadyCaught {
		fmt.Println("Pokemon already caught!")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	data := getPokemon(c, name)

	if data.BaseExperience == 0 {
		return errors.New("Pokemon not found")
	}

	var randomNumber int
	if data.BaseExperience < 50 {
		randomNumber = 100
	} else if data.BaseExperience < 100 {
		randomNumber = 125
	} else {
		randomNumber = 140
	}

	if randomNumber < data.BaseExperience {
		randomNumber += data.BaseExperience / 10
	}

	caughtNumber := rand.Intn(randomNumber + 1)
	outputMsg := fmt.Sprintf("%s escaped!", name)
	if caughtNumber > data.BaseExperience {
		outputMsg = fmt.Sprintf("%s was caught!", name)
		c.pokemon[name] = data
	}
	fmt.Println(outputMsg)
	return nil
}

func commandInspect(c *Config, name string) error {
	if name == "" {
		return errors.New("Pokemon name not provided")
	}

	pokemon, alreadyCaught := c.pokemon[name]
	if alreadyCaught == false {
		fmt.Println("Pokemon has not been caught!")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("\t- %s\n", v.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *Config, _ string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.pokemon {
		fmt.Printf("\t- %s\n", v.Name)
	}
	return nil
}
