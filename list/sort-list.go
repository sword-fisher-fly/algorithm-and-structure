package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func sortListHelper(head *ListNode, tail *ListNode) *ListNode {
//     if head == nil  {
//         return head
//     }

//     if head.Next == tail {
//         head.Next = nil
//         return head
//     }

//     slow, fast := head, head
//     for fast != tail {
//         slow = slow.Next
//         fast = fast.Next
//         if fast != tail {
//             fast = fast.Next
//         }
//     }

//     mid := slow
//     left := sortListHelper(head, mid)
//     right := sortListHelper(mid, tail)
//     return mergeSortedList(left, right)
// }

// func mergeSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
//     if l1 == nil {
//         return l2
//     }

//     if l2 == nil {
//         return l1
//     }

//     dummyNode := &ListNode{}
//     temp := dummyNode

//     for l1 != nil  && l2 != nil {
//         if l1.Val <= l2.Val {
//             temp.Next = l1
//             l1 = l1.Next
//         } else {
//             temp.Next = l2
//             l2 = l2.Next
//         }
//         temp = temp.Next
//     }

//     if l1 != nil {
//         temp.Next = l1
//     } else {
//         temp.Next = l2
//     }

//     return dummyNode.Next
// }

// func sortList(head *ListNode) *ListNode {
//     return sortListHelper(head, nil)
// }

// 2)

func mergeSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	temp := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}

	return dummyNode.Next
}

func sortListHelperII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := pre.Next
	pre.Next = nil
	left := sortListHelperII(head)
	right := sortListHelperII(mid)

	return mergeSortedList(left, right)
}

func SortListII(head *ListNode) *ListNode {
	return sortListHelperII(head)
}
