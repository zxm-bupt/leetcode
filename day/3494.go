package day

func minTime(skill []int, mana []int) int64 {
	startTime := int64(0)
	completeTime := make([]int64, len(skill))
	for _, v := range mana {
		preSum := make([]int64, len(skill))
		preSum[0] = int64(skill[0]) * int64(v)
		for i := 1; i < len(preSum); i++ {
			preSum[i] = preSum[i-1] + int64(skill[i])*int64(v)
			startTime = max(startTime, completeTime[i]-preSum[i-1])
		}
		for i := 0; i < len(skill); i++ {
			completeTime[i] = startTime + preSum[i]
		}
		startTime = completeTime[0]
	}
	return completeTime[len(skill)-1]
}
