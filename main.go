package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	res := day.RemoveAnagrams(words)
	fmt.Print(res)
}
