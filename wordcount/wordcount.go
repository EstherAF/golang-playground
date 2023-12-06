package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	wordsMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count := wordsMap[word]
		wordsMap[word] = count + 1
	}
	return wordsMap
}
