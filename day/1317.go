package day

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		if !strings.Contains(strconv.Itoa(a), "0") && !strings.Contains(strconv.Itoa(n-a), "0") {
			return []int{a, n - a}
		}
	}
	return []int{}
}
