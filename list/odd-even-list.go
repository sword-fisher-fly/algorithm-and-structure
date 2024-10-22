package list

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddHead := &ListNode{Next: nil, Val: 0}
	odd := oddHead
	evenHead := &ListNode{Next: nil, Val: 0}
	even := evenHead

	count := 1
	for head != nil {
		if count%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		count++
	}

	even.Next = nil
	odd.Next = evenHead.Next

	return oddHead.Next
}
