package day

func longestSubarray(nums []int) int {
	maxValue := nums[0]
	res, cnt := 0, 0
	for _, num := range nums {
		if num > maxValue {
			res, cnt = 1, 1
			maxValue = num
		} else if num == maxValue {
			cnt++
		} else {
			cnt = 0
		}
		res = max(res, cnt)
	}
	return res

}
func LongestSubarray(nums []int) int {
	return longestSubarray(nums)
}
