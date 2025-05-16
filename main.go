package main

import "leetcode/day"

func main() {
	// Example usage
	words := []string{"c"}
	groups := []int{0}

	// Call the function
	result := day.GetLongestSubsequence2(words, groups)
	// Print the result
	for _, word := range result {
		println(word)
	}
}
