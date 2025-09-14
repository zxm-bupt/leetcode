package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	res := day.Spellchecker(wordlist, queries)
	fmt.Println(res)
}
