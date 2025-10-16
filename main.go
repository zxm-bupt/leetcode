package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{0, -3}
	value := 4
	res := day.FindSmallestInteger(nums, value)
	fmt.Print(res)
}
