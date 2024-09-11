package list

type Node struct {
	Value int64
	Next  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func NewNode(value int64) *Node {
	return &Node{Value: value}
}

func NewList() *List {
	return &List{}
}

func (l *List) InitializeFromArray(nums []int64) {
	for _, v := range nums {
		l.Add(v)
	}
}

func (l *List) Add(value int64) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

func (l *List) AddNode(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *List) RemoveNode(n *Node) {
	if l.Head == nil {
		return
	}

	preNode, curNode := l.Head, l.Head
	for curNode != nil {
		if curNode == n {
			preNode.Next = curNode.Next

			if n == l.Tail {
				l.Tail = preNode
			}

			return
		}

		preNode = curNode
		curNode = curNode.Next
	}
}

func (l *List) Remove(val int64) {
	dummyNode := &Node{Next: l.Head}

	curNode := dummyNode
	for curNode.Next != nil {
		if curNode.Next.Value == val {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}

	l.Head = dummyNode.Next
	l.Tail = curNode
}

func (l *List) Reverse() {
	var pre, cur *Node
	tmpTail := l.Tail

	l.Tail = l.Head
	for cur = l.Head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	l.Head = tmpTail
}

func (l *List) ToArray() []int64 {
	var arr []int64
	curNode := l.Head
	for curNode != nil {
		arr = append(arr, curNode.Value)
		curNode = curNode.Next
	}

	return arr
}
