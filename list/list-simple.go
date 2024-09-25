package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromArray(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}

	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return head
}

func ConcatenateList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	cur := head1

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = head2

	return head1
}

func ToArray(head *ListNode) []int {
	res := []int{}
	cur := head

	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

func SwapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	cur := dummyHead

	// 1 -> 2 -> 3 -> 4 -> 5
	// 1) 2 -> 1 -> 3 -> 4 -> 5
	// cur.Next =1
	// cur.Next.Next =2
	// cur.Next.Next.Next = 3

	// tmp1 = 1
	// tmp2 = 3

	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next  // dummyHead -> 2
		cur.Next.Next = tmp1      // 2 -> 1
		cur.Next.Next.Next = tmp2 // 1 -> 3

		cur = cur.Next.Next // cur -> 1
	}

	return dummyHead.Next
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	// 1 -> 2 -> 3 -> 4 -> nil
	// loop 1
	// 1) next = 2
	// 2) 1 -> nil
	// 3) pre = 1
	// 4) cur = 2
	// loop 2
	// 1) next = 3
	// 2) 2 -> 1
	// 3) pre = 2
	// 4) cur = 3
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func ReverseListWithDummyNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{Next: nil}

	// 1->2->3->4-> nil
	curNode := head // cur -> 1
	// loop1
	// next -> 2
	// 1 -> nil
	// dummy.Next -> 1
	// cur -> 2
	// loop2
	// next -> 3
	// 2 -> 1
	// dummy.Next -> 2
	// cur -> 4
	for curNode != nil {
		next := curNode.Next //
		curNode.Next = dummyNode.Next
		dummyNode.Next = curNode
		curNode = next
	}

	return dummyNode.Next
}

func reverseListRecursive(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	next := cur.Next
	cur.Next = pre

	return reverseListRecursive(cur, next)
}
func ReverseListRecursive(head *ListNode) *ListNode {
	var pre *ListNode

	cur := head

	return reverseListRecursive(pre, cur)
}

func ReverseListRecursiveII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1->2->3->4->nil
	// newHead -> (4->3->2...)
	// head.Next -> 2 => head.Next.Next -> 1 => 4->3->2->1
	// head.Next -> nil => 4->3->2->1->nil

	newHead := ReverseListRecursiveII(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
