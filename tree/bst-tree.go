package tree

import (
	"fmt"
	"strings"
)

// Create
// Insert
// Find
// Delete // TODO

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

func NewBSTreeFromLevelArray(arr []any) *BSTree {
	if len(arr) == 0 {
		return &BSTree{root: nil}
	}

	treeNodeList := make([]*TreeNode, 0, len(arr))
	// leftChild: 2*i rightChild: 2*i+1
	for i := 0; i < len(arr); i++ {
		if arr[i] == nil {
			treeNodeList = append(treeNodeList, nil)
			continue
		}

		treeNodeList = append(treeNodeList, NewTreeNode(arr[i]))
	}

	for i := 0; i < len(arr)/2; i++ {
		if treeNodeList[i] == nil {
			continue
		}

		if 2*i+1 < len(arr) && treeNodeList[2*i+1] != nil {
			treeNodeList[i].Left = treeNodeList[2*i+1]
		}

		if 2*i+2 < len(arr) && treeNodeList[2*i+2] != nil {
			treeNodeList[i].Right = treeNodeList[2*i+2]
		}
	}

	return &BSTree{root: treeNodeList[0]}
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

// Example:
//
//	 2
//	/ \
//
// 1  3
// Ouput:
// 2 -> 1
// 2 -> 3
// 1 -> nil
// 3 -> nil
func (t *BSTree) PrettyPrint() {
	if t.Empty() {
		fmt.Println("Empty Tree")
		return
	}

	res := strings.Builder{}
	queue := []*TreeNode{t.root}
	level := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		res.WriteString(fmt.Sprintf("------------  level %d   ------\n", level))
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				res.WriteString(fmt.Sprintf("%d-leftChild->%d\n", node.Val, node.Left.Val))
				queue = append(queue, node.Left)
			} else {
				res.WriteString(fmt.Sprintf("%d-leftChild->nil\n", node.Val))

			}

			if node.Right != nil {
				res.WriteString(fmt.Sprintf("%d-rightChild->%d\n", node.Val, node.Right.Val))
				queue = append(queue, node.Right)
			} else {
				res.WriteString(fmt.Sprintf("%d-rightChild->nil\n", node.Val))
			}
		}
		level++
	}

	fmt.Println(res.String())
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
	var deleteNode func(root *TreeNode, val any) *TreeNode
	deleteNode = func(root *TreeNode, val any) *TreeNode {
		if root == nil {
			return nil
		}

		// TODO: core logic
		// 1) 左节点、右节点为空，直接删除
		// 2）左节点为空，右节点不为空
		// 3）左节点不为空，右节点为空
		// 4）左节点不为空，右节点不为空
		if val.(int) == root.Val.(int) {
			if root.Left == nil && root.Right == nil {
				return nil
			} else if root.Left != nil && root.Right == nil {
				return root.Left
			} else if root.Left == nil && root.Right != nil {
				return root.Right
			} else {
				// TODO: find the successor node
				// 1) find the right node
				// 2) find the left node
				// 3) find the left node's right node
				leftNode := root.Left
				curNode := root.Right
				for curNode.Left != nil {
					curNode = curNode.Left
				}

				curNode.Left = leftNode
				return root.Right
			}
		}

		if root.Val.(int) < val.(int) {
			root.Right = deleteNode(root.Right, val)
		} else {
			root.Left = deleteNode(root.Left, val)
		}

		return root
	}

	t.root = deleteNode(t.root, val)
}

func deleteNode(root *TreeNode, val any) *TreeNode {
	if root == nil {
		return root
	}

	// 删除某个节点
	var deleteOneNode func(target *TreeNode) *TreeNode
	deleteOneNode = func(target *TreeNode) *TreeNode {
		if target == nil {
			fmt.Println("target is nil")
			return target
		}

		if target.Right == nil {
			return target.Left
		}

		curNode := target.Right
		for curNode.Left != nil {
			curNode = curNode.Left
		}

		curNode.Left = target.Left
		return target.Right
	}

	curNode := root
	var preNode *TreeNode
	for curNode != nil {
		if curNode.Val.(int) == val.(int) {
			break
		}

		preNode = curNode
		if curNode.Val.(int) < val.(int) {
			curNode = curNode.Right
		} else {
			curNode = curNode.Left
		}
	}

	if preNode == nil {
		return deleteOneNode(curNode)
	}

	if preNode.Left != nil && preNode.Left.Val.(int) == val.(int) {
		preNode.Left = deleteOneNode(curNode)
	}

	if preNode.Right != nil && preNode.Right.Val.(int) == val.(int) {
		preNode.Right = deleteOneNode(curNode)
	}

	return root
}
func (t *BSTree) DeleteByRecursive(val any) {
	t.root = deleteNode(t.root, val)
	return
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
