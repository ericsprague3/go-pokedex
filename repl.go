package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		firstArg := cleaned[0]
		var secondArg string
		if len(cleaned) == 2 {
			secondArg = cleaned[1]
		}

		cmd, exists := cfg.commands[firstArg]
		if exists {
			err := cmd.callback(cfg, secondArg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

}
