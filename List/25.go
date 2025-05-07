package List

func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	current := dummy
	if head == nil {
		return nil
	}
	count := 0
	for head != nil {
		count++
		if count == k {
			temp := head.Next
			head.Next = nil
			head = temp
			current.Next = ResverseList(current.Next)
			for current.Next != nil {
				current = current.Next
			}
			current.Next = head
			count = 0
		} else {
			head = head.Next
		}
	}

	return dummy.Next
}
