package day

import "sort"

func PartitionArray(nums []int, k int) int {
	sort.Ints(nums)
	count := 1
	temp := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i]-temp > k {
			temp = nums[i]
			count++
		}
	}
	return count
}
