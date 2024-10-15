package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO:https://leetcode.cn/problems/insert-into-a-sorted-circular-linked-list/
// 给定循环升序列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的

func InsertSortedCircleList(head *ListNode, x int) *ListNode {
	if head == nil {
		head = &ListNode{Val: x}
		head.Next = head
		return head
	}

	// cur := head

	return head

}
