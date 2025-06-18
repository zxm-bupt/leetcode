package day

import "leetcode/sort"

func DivideArray(nums []int, k int) [][]int {
	res := make([][]int, 0)
	sort.QuicSort(nums)
	count := 0
	temp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		temp = append(temp, nums[i])
		count++
		if count == 3 {
			if nums[i]-temp[0] > k {
				return make([][]int, 0)
			}
			res = append(res, temp)
			temp = make([]int, 0)
			count = 0
		}
	}
	return res
}
