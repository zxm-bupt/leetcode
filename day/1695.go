package day

func maximumUniqueSubarray(nums []int) int {
	freSum := make([]int, len(nums)+1)
	freSum[1] = nums[0]
	startIndex := 0
	hash := make(map[int]int, len(nums))
	hash[nums[0]] = 0
	res := nums[0]
	for n := 1; n < len(nums); n++ {
		freSum[n+1] = freSum[n] + nums[n]
		var lastLocation int
		if location, ok := hash[nums[n]]; ok {
			lastLocation = location
		} else {
			lastLocation = -1
		}

		hash[nums[n]] = n
		startIndex = max(startIndex, lastLocation+1)
		res = max(res, freSum[n+1]-freSum[startIndex])
	}
	return res
}

func MaximumUniqueSubarray(nums []int) int {
	return maximumUniqueSubarray(nums)
}
