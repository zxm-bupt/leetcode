package day

import "sort"

func findLHS(nums []int) int {
	res := 0
	sort.Ints(nums)
	left, right := 0, 0
	for right < len(nums) {
		if nums[right]-nums[left] == 1 {
			res = max(res, right-left+1)
			right++
		} else if nums[right]-nums[left] > 1 {
			left++
		} else {
			right++
		}
	}
	return res
}
