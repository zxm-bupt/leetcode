package main

import "leetcode/List"

func main() {
	// Example usage
	head := &List.ListNode{Val: 1}
	List.ReorderList(head)
	for node := head; node != nil; node = node.Next {
		println(node.Val)
	}
	// Output: 1 -> 4 -> 2 -> 3
}
