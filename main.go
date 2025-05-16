package main

import (
	"fmt"
	"leetcode/List"
)

func main() {
	// Example usage
	l1 := &List.ListNode{Val: 2}
	l1.Next = &List.ListNode{Val: 4}
	l1.Next.Next = &List.ListNode{Val: 3}

	l2 := &List.ListNode{Val: 5}
	l2.Next = &List.ListNode{Val: 6}
	l2.Next.Next = &List.ListNode{Val: 4}

	res := List.AddTwoNumbers(l1, l2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}
