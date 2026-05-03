package main

import (
	"sort"
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

func sortMapKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
