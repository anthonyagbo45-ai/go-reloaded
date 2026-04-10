package main

import (
	"strings"
)

func Capitalize(word string) string {
	words := strings.Fields(word)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)"&& i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}

func Lower(word string) string {
	words := strings.Fields(word)

	for r := 0; r < len(words); r++ {
		if words[r] == "(low)" && r > 0 {
			words[r-1] = strings.ToLower(words[r-1])
			words[r] = ""
		}
	}
	return strings.Join(words, " ")
}
 