package main

import "leetcode/day"

func main() {
	nums := []int{3, 1, 2, 4, 5}

	k := 3

	res := day.MaxSubsequence(nums, k)
	for _, v := range res {
		println(v)
	}
}
