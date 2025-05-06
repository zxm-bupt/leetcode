package List

type ListNode struct {
	Val  int
	Next *ListNode
}

func resverseList(head *ListNode) *ListNode {
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
