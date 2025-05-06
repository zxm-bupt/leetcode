package List

// my func
func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		return nil
	}
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
	current.Next = nil
	return dummy.Next
}

// leetcode func
func DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	if current != nil {
		if current.Next.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}
