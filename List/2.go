package List

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		newNode := &ListNode{Val: carry}
		if l1 != nil {
			newNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			newNode.Val += l2.Val
			l2 = l2.Next
		}
		carry = newNode.Val / 10
		newNode.Val %= 10
		current.Next = newNode
		current = current.Next
	}
	return dummy.Next
}
