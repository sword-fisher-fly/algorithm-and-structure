package list

func DeleteDuplicatesRemainOnlyExistOnce(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	cur := dummyNode

	// 1->1->2->3->3->nil
	// 2->nil
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyNode.Next
}

func RemoveDuplicatesRemainOnlyOne(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
