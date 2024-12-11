package list

//https://leetcode.cn/problems/reverse-linked-list-ii/?envType=study-plan-v2&envId=top-interview-150
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	cur := dummyNode
	var pre *ListNode
	for i := 0; i < left; i++ {
		pre = cur
		cur = cur.Next
	}

	firstStart := cur

	var lastEnd *ListNode
	for i := 0; i < right-left; i++ {
		cur = cur.Next
	}
	lastEnd = cur

	// fmt.Printf("left=%d, val=%d, right=%d,val=%d\n", left, firstStart.Val, right, lastEnd.Val)

	reverseList := func(first, last *ListNode) *ListNode {
		pre := last
		cur := first
		for cur != last {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		return pre
	}

	res := reverseList(firstStart, lastEnd.Next)

	pre.Next = res

	return dummyNode.Next
}
