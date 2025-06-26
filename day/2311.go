package day

import "math"

func LongestSubsequence(s string, k int) int {

	dp := make([]int, len(s)+1)
	dp[len(s)] = 0
	current := 0.0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = dp[i+1] + 1
		} else {
			current += math.Pow(2, float64(dp[i+1]))
			if current > float64(k) {
				dp[i] = dp[i+1]
				current -= math.Pow(2, float64(dp[i+1]))
			} else {
				dp[i] = dp[i+1] + 1
			}
		}

	}
	return dp[0]

}
