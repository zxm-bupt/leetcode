package day

func MaxFreeTime(eventTime int, startTime []int, endTime []int) int {
	return maxFreeTime(eventTime, startTime, endTime)
}
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	q := make([]bool, len(startTime))
	maxLeftRest := startTime[0]
	maxRightRest := eventTime - endTime[len(endTime)-1]
	for i := 1; i < len(startTime); i++ {
		during := endTime[i] - startTime[i]
		if maxLeftRest >= during {
			q[i] = true
		}
		left := startTime[i] - endTime[i-1]
		maxLeftRest = max(maxLeftRest, left)
	}
	for j := len(endTime) - 2; j >= 0; j-- {
		during := endTime[j] - startTime[j]
		if !q[j] && maxRightRest >= during {
			q[j] = true
		}
		right := startTime[j+1] - endTime[j]
		maxRightRest = max(maxRightRest, right)
	}
	maxTime := 0
	for i := 0; i < len(startTime); i++ {
		during := endTime[i] - startTime[i]
		var start, end, rest int
		if i == 0 {
			start = 0
		} else {
			start = endTime[i-1]
		}
		if i == len(startTime)-1 {
			end = eventTime
		} else {
			end = startTime[i+1]
		}
		if q[i] {
			rest = end - start
		} else {
			rest = end - start - during
		}
		maxTime = max(maxTime, rest)
	}
	return maxTime
}
