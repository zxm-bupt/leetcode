package main

import "leetcode/day"

func main() {
	nums := []int{2, 4, 2, 2, 5, 2}
	k := 2
	result := day.DivideArray(nums, k)
	for _, group := range result {
		for _, num := range group {
			print(num, " ")
		}
		println()
	}
}
