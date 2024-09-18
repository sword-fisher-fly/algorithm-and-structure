package list

// list.List
type DoubleLinkedList struct {
	Head, Tail *ListNode
}

func InitializeDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) Empty() bool {
	return l.Head == nil && l.Tail == nil
}
