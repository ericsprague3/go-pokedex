package main

import (
	"strings"
)

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	words := strings.Fields(strings.ToLower(text))
	return words
}
