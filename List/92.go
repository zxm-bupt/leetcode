package List

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	leftNode := pre.Next
	leftTail := pre

	//reverse
	pre = pre.Next
	var current *ListNode
	current = nil
	for i := left; i < right; i++ {
		next := pre.Next
		pre.Next = current
		current = pre
		pre = next
	}
	next := pre.Next
	pre.Next = current
	leftTail.Next = pre
	leftNode.Next = next
	return dummy.Next
}
