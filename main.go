package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	points := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	res := day.LargestTriangleArea(points)
	fmt.Println(res)
}
