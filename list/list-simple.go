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
	if head == nil || head.Next == nil {
		return head
	}

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

		cur = cur.Next.Next
	}

	return dummyHead.Next
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
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

	newHead := ReverseListRecursiveII(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
