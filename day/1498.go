package day


func numSubseq(nums []int, target int) int {
	pow := make([]int, len(nums))
	pow[0] = 1
	for i := 1; i < len(nums); i++ {
		pow[i] = pow[i-1] * 2 % (1e9 + 7)
	}
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	res := 0
	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			res = (res + pow[right-left]) % (1e9 + 7)
			left++
		}
	}

	return res
}