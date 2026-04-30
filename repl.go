package main

import (
	"strings"
)

func cleanInput(text string) (cleanText []string) {
	splitText := strings.Split(strings.ToLower(text), " ")
	for _, word := range splitText {
		if word != "" {
			cleanText = append(cleanText, word)
		}
	}
	return cleanText
}
