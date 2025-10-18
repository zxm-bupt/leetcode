package day

import (
	"math"
	"sort"
)

func MaxDistinctElements(nums []int, k int) int {
	return maxDistinctElements(nums, k)
}

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	cnt := 0
	prev := math.MinInt32

	for _, num := range nums {
		curr := min(max(num-k, prev+1), num+k)
		if curr > prev {
			cnt++
			prev = curr
		}
	}
	return cnt
}
