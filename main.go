package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	nums := []int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825}
	res := day.ReplaceNonCoprimes(nums)
	fmt.Println(res)
}
