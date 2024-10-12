package offer

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func NewBSTreeFromSortedArray(arr []interface{}) *BSTree {
	tree := &BSTree{}
	tree.root = tree.buildFromSortedArray(arr, 0, len(arr)-1)
	return tree
}

// TODO:
func NewBSTreeFromLevelOrderArray(arr []interface{}) *BSTree {
	tree := &BSTree{}
	// tree.root = tree.buildFromLevelOrderArray(arr, 0)
	return tree
}

func (t *BSTree) buildFromSortedArray(arr []interface{}, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &TreeNode{Val: arr[mid]}
	node.Left = t.buildFromSortedArray(arr, start, mid-1)
	node.Right = t.buildFromSortedArray(arr, mid+1, end)
	return node
}

func (t *BSTree) ContructIterator() *BSTIterator {
	iterator := &BSTIterator{}
	iterator.stack = []*TreeNode{}

	root := t.root
	for root != nil {
		iterator.stack = append(iterator.stack, root)
		root = root.Left
	}

	return iterator
}

func (i *BSTIterator) HasNext() bool {
	return len(i.stack) > 0
}

func (i *BSTIterator) Next() interface{} {
	curNode := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]

	node := curNode.Right
	for node != nil {
		i.stack = append(i.stack, node)
		node = node.Left
	}

	return curNode.Val
}
