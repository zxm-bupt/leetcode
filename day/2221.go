package day

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp := make([]int, len(nums)-1)
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		temp[count] = (nums[i] + nums[i+1]) % 10
		count++
	}
	return triangularSum(temp)
}
