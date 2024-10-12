package list

// 归并排序的思想

// 1) 获取链表的中间节点, 分割为两个子链表
// 2) 递归调用链表排序函数, 使两个子链表排序
// 3) 合并两个有序链表

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1,2,3,4,5,6,7,8,9,10
	// 1) slow: 1 fast: 2
	// 2) slow: 2 fast: 4
	// 3) slow: 3 fast: 6
	// 4) slow: 4 fast: 8
	// 5) slow: 5 fast: 10
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	l2 := slow.Next
	slow.Next = nil

	l1 := SortList(head)
	l2 = SortList(l2)

	return MergeSortedListByIteration(l1, l2)
}
