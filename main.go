package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{311155, 311155, 311155, 311155, 311155, 311155, 311155, 311155, 201191, 311155}
	res := day.LongestSubarray(nums)
	fmt.Println(res)
}
