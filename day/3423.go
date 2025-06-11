package day

import "math"

func maxAdjacentDistance(nums []int) int {
	result := 0
	for index, value := range nums {
		nextIndex := (index + 1) % len(nums)
		nextValue := nums[nextIndex]
		distance := float64(value - nextValue)
		result = max(result, int(math.Abs(distance)))
	}

	return result
}
