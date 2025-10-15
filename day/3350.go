package day

func maxIncreasingSubarrays(nums []int) int {
	cnt := 1
	precnt := 0
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			precnt = cnt
			cnt = 1
		} else {
			cnt++
		}
		ans = max(ans, min(precnt, cnt), cnt/2)
	}

	return ans
}
