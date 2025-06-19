package main

import "leetcode/day"

func main() {
	nums := []int{2, 4, 2, 2, 5, 2}
	k := 2
	result := day.PartitionArray(nums, k)
	println(result) // Output: 2
}
