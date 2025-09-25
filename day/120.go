package day

import "math"

func MinimumTotal(triangle [][]int) int {
	return minimumTotal(triangle)
}

func minimumTotal(triangle [][]int) int {
	res := math.MaxInt32
	var getValue func(int, int, int)
	getValue = func(length int, value int, index int) {
		if length == len(triangle) {
			res = min(value, res)
			return
		}
		value = value + triangle[length][index]
		getValue(length+1, value, index)
		getValue(length+1, value, index+1)
	}
	getValue(0, 0, 0)
	return res
}

func minimumTotal_2(triangle [][]int) int {
	type location struct {
		row int
		col int
	}
	hashValue := make(map[location]int)
	for i := len(triangle) - 1; i >= 0; i-- {
		for index, value := range triangle[i] {
			leftValue := 0
			rightValue := math.MaxInt32
			if left, ok := hashValue[location{i + 1, index}]; ok {
				leftValue = left
			}
			if right, ok := hashValue[location{i + 1, index + 1}]; ok {
				rightValue = right
			}
			hashValue[location{i, index}] = value + min(rightValue, leftValue)

		}

	}
	return hashValue[location{0, 0}]

}
