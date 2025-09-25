package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := day.MinimumTotal(triangle)
	fmt.Println(res)
}
