package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{558, 508, 782, 32, 187, 103, 370, 607, 619, 267, 984, 10}
	res := day.MaximumUniqueSubarray(nums)
	fmt.Println(res)
}
