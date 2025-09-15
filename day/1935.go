package day

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	res := 0
	for _, word := range words {
		if canBeTypedWord(word, brokenLetters) {
			res++
		}
	}
	return res
}

func canBeTypedWord(word string, brokenLetters string) bool {
	for _, char := range brokenLetters {
		if strings.ContainsRune(word, char) {
			return false
		}
	}
	return true
}
