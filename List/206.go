package List

func ResverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	current := head.Next
	post := head
	head.Next = nil
	for current != nil {
		temp := current.Next
		current.Next = post
		post = current
		current = temp
	}

	return post
}
