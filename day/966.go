package day

import (
	"strings"
)

func Spellchecker(wordlist []string, queries []string) []string {
	return spellchecker(wordlist, queries)
}

func spellchecker(wordlist []string, queries []string) []string {
	res := make([]string, len(queries))
	for i, query := range queries {
		ans := spellcheckhelper(wordlist, query)
		res[i] = ans
	}
	return res
}

func spellcheckhelper(wordlist []string, query string) string {
	for _, world := range wordlist {
		if world == query {
			return world
		}
	}
	for _, world := range wordlist {
		if strings.EqualFold(world, query) {
			return world
		}
	}
	for _, world := range wordlist {
		if vowelchecker(world, query) {
			return world
		}
	}
	return ""
}

func vowelchecker(word string, query string) bool {
	f := []byte(word)
	s := []byte(query)
	if len(f) != len(s) {
		return false
	}
	for i := 0; i < len(f); i++ {
		if f[i] != s[i] && !(isvowel(f[i]) && isvowel(s[i])) && (f[i]+32 != s[i] && f[i]-32 != s[i]) {
			return false
		}
	}
	return true

}

func isvowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
