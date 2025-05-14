package main

import "leetcode/day"

func main() {
	// Example usage
	s := "abcyy"
	t := 2
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	result := day.LengthAfterTransformations2(s, t, nums)
	println(result)
}
