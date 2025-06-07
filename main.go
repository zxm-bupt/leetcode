package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{1, 2, 1}
	k := 3
	edges := [][]int{{0, 1}, {0, 2}}
	fmt.Println(day.MaximumValueSum(nums, k, edges))

}
