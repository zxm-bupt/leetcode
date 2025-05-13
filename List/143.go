package List

func ReorderList(head *ListNode) {
	mid, _ := findMid(head)
	temp := mid.Next
	mid.Next = nil
	if temp == nil {
		return
	}
	mid = temp
	// reverse the second half
	second := ResverseList(mid)
	// merge two lists
	first := head
	for second != nil {
		temp1 := first.Next
		temp2 := second.Next
		first.Next = second
		second.Next = temp1
		first = temp1
		second = temp2
	}
}

func findMid(head *ListNode) (*ListNode, bool) {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	symbol := false
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		} else {
			symbol = true
		}
	}

	return slow, symbol
}
