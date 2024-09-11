package list

func IntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB

	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
