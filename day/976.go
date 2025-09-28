package day

import "sort"

func isTri(a, b, c int) bool {
	return a+b > c && b-a < c
}

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := len(nums) - 1; i > 1; i-- {

		if isTri(nums[i-2], nums[i-1], nums[i]) {
			return nums[i-2] + nums[i-1] + nums[i]
		}

	}

	return 0
}
