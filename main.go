package main

import (
	"fmt"
	"leetcode/List"
)

func main() {
	node := &List.ListNode{Val: 1}
	node.Next = &List.ListNode{Val: 2}
	node.Next.Next = &List.ListNode{Val: 3}
	node.Next.Next.Next = &List.ListNode{Val: 4}
	node.Next.Next.Next.Next = &List.ListNode{Val: 5}

	fmt.Println(List.ReverseBetween(node, 2, 4))
}
