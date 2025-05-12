package List

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		pA = pA.Next
		pB = pB.Next
		if pA == nil && pB == nil {
			return nil
		}
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
	}
	return pA
}
