package main

import (
	"fmt"
	"leetcode/sort"
)

func main() {
	// Example usage
	arr := []int{10, 7, 8, 9, 1, 5}
	sort.QuicSort(arr)
	fmt.Println("Sorted array:", arr)
}
