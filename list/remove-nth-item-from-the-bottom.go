package list

func RemoveNthItemFromTheBottomII(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	slow, fast := dummyHead, dummyHead

	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	// n: 1
	// fastNode: 1
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	// 保护
	if fast == nil {
		return head
	}

	fast = fast.Next // fastNode -> 2

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}

// TODO:
func RemoveNthItemFromTheBottomByRecursive(head *ListNode, n int) *ListNode {
	return nil
}
