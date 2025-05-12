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
	r := List.MergeTwoSeqList(&list, &list2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
