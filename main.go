package main

import (
	"fmt"
	"leetcode/List"
)

func main() {
	// Example usage
	head := &List.ListNode{Val: 4}
	head.Next = &List.ListNode{Val: 2}
	head.Next.Next = &List.ListNode{Val: 1}
	head.Next.Next.Next = &List.ListNode{Val: 3}
	sortedHead := List.SortList(head)
	// Print sorted list
	for sortedHead != nil {
		fmt.Println(sortedHead.Val)
		sortedHead = sortedHead.Next
	}
}
