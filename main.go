package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	powers := []int{7, 1, 6, 3}
	res := day.MaximumTotalDamage(powers)
	fmt.Print(res)
}
