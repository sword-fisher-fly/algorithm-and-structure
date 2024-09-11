package list

// func RemoveNthItemFromTheBottom(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	listLength := 0
// 	curNode := head

// 	for curNode != nil {
// 		listLength++
// 		curNode = curNode.Next
// 	}

// 	if n > listLength {
// 		return head
// 	}

// 	// 1 -> 2 -> 3 -> 4 -> 5
// 	// n: 1
// 	// fastNode: 2
// 	fastNode := head
// 	for i := 0; i < n-1 && fastNode != nil; i++ {
// 		fastNode = fastNode.Next
// 	}

// 	fastNode = fastNode.Next

// 	curNode = head
// 	for fastNode != nil {
// 		curNode = curNode.Next
// 		fastNode = fastNode.Next
// 	}

// 	if curNode.Next == nil {
// 		return head.Next
// 	}

// 	curNode.Next = curNode.Next.Next

// 	return head
// }

// 保证n小于等于链表长度
func RemoveNthItemFromTheBottomII(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	slow, fast := dummyHead, dummyHead

	// 1 -> 2 -> 3 -> 4 -> 5
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
