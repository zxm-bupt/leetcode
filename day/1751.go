package day

import "sort"

func MaxValue(events [][]int, k int) int {
	return maxValue(events, k)
}

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([][]int, len(events)+1)
	for i := 0; i < len(events)+1; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i < len(events)+1; i++ {
		startTime := events[i-1][0]
		left, right := 0, len(events)-1
		for left <= right {
			mid := (left + right) / 2
			if events[mid][1] >= startTime {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		//left := sort.Search(i, func(j int) bool { return events[j][1] >= events[i][0] })
		for j := 1; j < k+1; j++ {
			dp[i][j] = max(dp[i-1][j], dp[left][j-1]+events[i-1][2])

		}
	}

	return dp[len(events)][k]
}
