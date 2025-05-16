package day

func GetWordsInLongestSubsequence3(words []string, groups []int) []string {
	dp := make([]int, len(words))
	prev := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		prev[i] = -1
	}
	max := 0
	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if groups[i] == groups[j] || getHanmingDistance(words[i], words[j]) != 1 {
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

func getHanmingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}
