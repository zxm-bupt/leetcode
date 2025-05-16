package day

func GetLongestSubsequence(words []string, groups []int) []string {
	index := []int{0}
	flag := groups[0]
	for i := 1; i < len(groups); i++ {
		if groups[i] == flag {
			continue
		}
		index = append(index, i)
		flag = groups[i]
	}
	res := []string{}
	for i := 0; i < len(index); i++ {
		location := index[i]
		s := words[location]
		res = append(res, s)
	}
	return res
}

// 动态规划算法
func GetLongestSubsequence2(words []string, groups []int) []string {
	dp := make([]int, len(words))
	prev := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		prev[i] = -1
	}
	max := 0
	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if groups[i] == groups[j] {
				continue
			}
			if dp[i] < dp[j]+1 {
				prev[i] = j
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > dp[max] {
			max = i
		}
	}
	res := []string{}
	for max != -1 {
		res = append(res, words[max])
		max = prev[max]
	}
	// 反转结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
