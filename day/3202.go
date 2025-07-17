package day

func maximumLength2(nums []int, k int) int {
	dp := make([][]int, k)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k)
	}
	res := 0
	for _, num := range nums {
		for i := 0; i < k; i++ {
			dp[i][num%k] = dp[num%k][i] + 1
			res = max(res, dp[i][num%k])
		}
	}
	return res
}
func MaximumLength2(nums []int, k int) int {
	return maximumLength2(nums, k)
}
