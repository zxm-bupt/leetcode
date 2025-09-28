package day

import (
	"sort"
)

func TriangleNumber(nums []int) int {
	return triangleNumber(nums)
}
func triangleNumber(nums []int) int {
	res := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			start := sort.Search(len(nums), func(a int) bool {
				if a <= j {
					return false
				} else {
					return nums[a] > nums[j]-nums[i]
				}
			})
			end := sort.Search(len(nums), func(a int) bool {

				if a <= j {
					return false
				} else {
					return nums[a] >= nums[j]+nums[i]
				}

			})
			res = res + end - start
		}

	}
	return res
}
