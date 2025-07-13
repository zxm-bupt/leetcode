package day

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	return getDecimalValue(head)
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res * 2
		res += head.Val
		head = head.Next
	}
	return res
}
