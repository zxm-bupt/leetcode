package day

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	sum := make([]int, len(startTime)+1)
	for i := 1; i < len(startTime)+1; i++ {
		time := endTime[i-1] - startTime[i-1]
		sum[i] = time + sum[i-1]
	}
	res := 0
	for i := k - 1; i < len(startTime); i++ {
		var left, right int
		if i == k-1 {
			left = 0
		} else {
			left = endTime[i-k]
		}
		if i == len(startTime)-1 {
			right = eventTime
		} else {
			right = startTime[i+1]
		}
		res = max(res, right-left-sum[i+1]+sum[i-k+1])
	}
	return res
}
