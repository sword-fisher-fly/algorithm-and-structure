package tree

// Create
// Insert
// Find
// Delete

type BSTree struct {
	root *TreeNode
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func buildBSTreeFromSortedArray(arr []any, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	root := &TreeNode{
		Val: arr[mid],
	}

	root.Left = buildBSTreeFromSortedArray(arr, left, mid-1)
	root.Right = buildBSTreeFromSortedArray(arr, mid+1, right)

	return root
}
func NewBSTreeFromSortedArray(arr []any) *BSTree {
	root := buildBSTreeFromSortedArray(arr, 0, len(arr)-1)

	return &BSTree{root: root}
}

func NewBSTreeFromSortedArrayByIteration(arr []any) *BSTree {
	if len(arr) == 0 {
		return &BSTree{}
	}

	root := &TreeNode{}
	nodeQueue := []*TreeNode{root}
	leftQueue := []int{0}
	rightQueue := []int{len(arr) - 1}
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		left := leftQueue[0]
		leftQueue = leftQueue[1:]

		right := rightQueue[0]
		rightQueue = rightQueue[1:]

		mid := left + (right-left)/2
		node.Val = arr[mid]

		if left <= mid-1 {
			node.Left = &TreeNode{}
			nodeQueue = append(nodeQueue, node.Left)
			leftQueue = append(leftQueue, left)
			rightQueue = append(rightQueue, mid-1)
		}

		if right >= mid+1 {
			node.Right = &TreeNode{}
			nodeQueue = append(nodeQueue, node.Right)
			leftQueue = append(leftQueue, mid+1)
			rightQueue = append(rightQueue, right)
		}
	}

	return &BSTree{root: root}
}

func (t *BSTree) Empty() bool {
	return t.root == nil
}

func (t *BSTree) Root() *TreeNode {
	return t.root
}

func (t *BSTree) PreOrder() []any {
	ret := []any{}

	PreOrder(t.root, &ret)

	return ret
}

func (t *BSTree) InOrder() []any {
	ret := []any{}

	InOrder(t.root, &ret)

	return ret
}

func (t *BSTree) PostOrder() []any {
	ret := []any{}

	PostOrder(t.root, &ret)

	return ret
}

func (t *BSTree) LevelTraversal() [][]any {
	return LevelTraversal(t.root)
}

// TODO: Add/Delete/Search
func (t *BSTree) SearchByRecursive(val any) *TreeNode {
	return searchByRecursive(t.root, val)
}

func (t *BSTree) Search(val any) *TreeNode {
	root := t.root
	for root != nil {
		if val.(int) < root.Val.(int) {
			root = root.Left
		} else if val.(int) > root.Val.(int) {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}

func (t *BSTree) Insert(val any) {
	if t.root == nil {
		node := NewTreeNode(val)
		t.root = node
		return
	}

	cur := t.root
	parent := t.root
	for cur != nil {
		parent = cur
		if val.(int) < cur.Val.(int) {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	node := NewTreeNode(val)

	if val.(int) < parent.Val.(int) {
		parent.Left = node
	} else {
		parent.Right = node
	}
}

func (t *BSTree) InsertByRecursive(val any) {
	t.root = insertByRecursive(t.root, val)
}

func (t *BSTree) Delete(val any) {

}

func searchByRecursive(root *TreeNode, val any) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val.(int) == val.(int) {
		return root
	}

	if val.(int) < root.Val.(int) {
		return searchByRecursive(root.Left, val)
	}

	return searchByRecursive(root.Right, val)
}

func insertByRecursive(root *TreeNode, val any) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}

	if val.(int) < root.Val.(int) {
		root.Left = insertByRecursive(root.Left, val)
	} else {
		root.Right = insertByRecursive(root.Right, val)
	}

	return root
}
