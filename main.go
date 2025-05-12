package main

import (
	"fmt"
	"leetcode/List"
)

func main() {
	list := List.ListNode{Val: 1}
	list.Next = &List.ListNode{Val: 2}
	list.Next.Next = &List.ListNode{Val: 3}
	list2 := List.ListNode{Val: 1}
	list2.Next = &List.ListNode{Val: 2}
	list2.Next.Next = &List.ListNode{Val: 3}
	list2.Next.Next.Next = list.Next
	fmt.Print(List.GetIntersectionNode(&list, &list2).Val) // 3
}
