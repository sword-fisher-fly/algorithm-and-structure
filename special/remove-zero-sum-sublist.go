package special

// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

// Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.

// After doing so, return the head of the final linked list.  You may return any such answer.

// (Note that in the examples below, all sequences are serializations of ListNode objects.)

// Example 1:

// ```text
// Input: head = [1,2,-3,3,1]
// Output: [3,1]
// Note: The answer [1,2,1] would also be accepted.
// ```
func RemoveZeroSumSublist(head *ListNode) *ListNode {
	headPre := &ListNode{
		Val:  2000000, // bigger than 1000*1000
		Next: head,
	}

	node := make(map[int]*ListNode, 1000)
	stack, top := make([]int, 1000), -1
	sum, cur := 0, headPre
	for cur != nil {
		sum += cur.Val
		if node[sum] == nil {
			node[sum] = cur
			top++
			stack[top] = sum
		} else {
			node[sum].Next = cur.Next
			for stack[top] != sum {
				node[stack[top]] = nil
				top--
			}
		}
		cur = cur.Next
	}

	return headPre.Next
}
