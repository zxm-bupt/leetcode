package main

import (
	"fmt"
	"leetcode/List"
)

func main() {
	node := &List.ListNode{
		Val: 1,
	}
	node.Next = &List.ListNode{
		Val: 2,
	}
	node.Next.Next = &List.ListNode{
		Val: 3,
	}
	node = List.ResverseList(node)
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}
