package main

import (
	"fmt"
	"leetcode/day"
)

func main() {
	n := 6
	delay := 2
	forget := 4
	res := day.PeopleAwareOfSecret(n, delay, forget)
	fmt.Println(res)
}
