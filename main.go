package main

import "leetcode/day"

func main() {
	events := [][]int{
		{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3},
	}

	result := day.MaxEvents(events)
	println(result) // Output: 4

}
