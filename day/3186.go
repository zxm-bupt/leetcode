package day

import "sort"

func MaximumTotalDamage(power []int) int64 {
	return maximumTotalDamage(power)
}

// 比较难
func maximumTotalDamage(power []int) int64 {
	count := make(map[int]int)
	res := int64(0)
	for _, v := range power {
		count[v]++
	}

	damage := make([]int, 0, len(count))

	for k, _ := range count {
		damage = append(damage, k)
	}

	sort.Ints(damage)

	dp := make([]int64, len(damage))
	//存储当前的mx
	j := 0
	mx := int64(0)
	for i := 0; i < len(damage); i++ {
		for ; damage[j] < damage[i]-2 && j < i; j++ {
			mx = max(mx, dp[j])
		}
		dp[i] = mx + int64(count[damage[i]])*int64(damage[i])
		res = max(res, dp[i])
	}

	return res

}
