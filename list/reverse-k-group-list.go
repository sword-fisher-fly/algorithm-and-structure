package list

// 1 -> 2 -> 3 -> 4 ...
// 1) first = 1
// 2) last = 4
// loop 1
// 1) next = 2
// 2) 1 -> 4
// 3) preNode = 1
// 4) first = 2
func reverseBetweenTwoNodes(first *ListNode, last *ListNode) *ListNode {
	preNode := last
	for first != last {
		next := first.Next
		first.Next = preNode
		preNode = first
		first = next
	}

	return preNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	// k = 3
	// node -> 4
	// 3 -> 2 -> 1 -> 4 -> 5 -> nil
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead := reverseBetweenTwoNodes(head, node)
	head.Next = ReverseKGroup(node, k)

	return newHead
}
