package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	success := 16
	res := day.SuccessfulPairs(spells, potions, int64(success))
	fmt.Print(res)
}
