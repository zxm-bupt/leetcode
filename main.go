package main

import "leetcode/day"

func main() {
	events := [][]int{
		//{87, 95, 42}, {3, 42, 37}, {20, 42, 100}, {53, 84, 80}, {10, 88, 38}, {25, 80, 57}, {18, 38, 33},

		{1, 2, 4}, {3, 4, 3}, {2, 3, 1},
	}

	result := day.MaxValue(events, 2)
	println(result) // Output: 4

}
