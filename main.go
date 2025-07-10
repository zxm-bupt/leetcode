package main

import "leetcode/day"

func main() {
	eventTime := 5
	startTime := []int{
		1, 3,
	}
	endTime := []int{
		2, 5,
	}
	result := day.MaxFreeTime(eventTime, startTime, endTime)
	println(result) // Output: 4

}
