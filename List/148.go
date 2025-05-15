package List

// 对于链表的归并排序，使用快慢指针找到中点，然后递归地对前后两部分进行排序，最后合并两个有序链表。
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid, _ := findMid(head)
	mid, mid.Next = mid.Next, nil
	left := SortList(head)
	right := SortList(mid)
	return MergeTwoSeqList(left, right)
}
