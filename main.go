package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	rains := []int{1, 2, 0, 1, 2}
	res := day.AvoidFlood(rains)
	fmt.Print(res)
}
