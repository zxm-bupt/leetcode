package List

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := head
	current := head
	for pre != nil {
		if pre.Val != current.Val {
			current.Next = pre
			current = pre
		}
		pre = pre.Next
	}
	return dummy.Next
}
