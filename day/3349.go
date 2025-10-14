package day

func HasIncreasingSubarrays(nums []int, k int) bool {
	return hasIncreasingSubarrays(nums, k)
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	count := 1
	precnt := 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			precnt = count
			count = 1
		} else {
			count++
		}
		ans = max(ans, min(precnt, count), count/2)
	}
	return ans >= k
}
