package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{-15, 9}
	k := 1
	res := day.HasIncreasingSubarrays(nums, k)
	fmt.Print(res)
}
