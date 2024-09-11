package tree

type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTreeNode(val any) *TreeNode {
	return &TreeNode{Val: val}
}

// Full binary tree
//
//	   1
//	  / \
//	 2   3
//	/\  / \
//
// 4  5 6 -1
// [1,2,3,4,5,6,-1]
// preOrder:  1 2 4 5 3 6
// inOrder:   4 2 5 1 6 3
// postOrder: 4 5 2 6 3 1
func NewTreeFromArray(arr []any) *Tree {
	if len(arr) == 0 {
		return &Tree{root: nil}
	}

	treeNodeList := make([]*TreeNode, 0, len(arr))
	// leftChild: 2*i rightChild: 2*i+1
	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 {
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

	return &Tree{root: treeNodeList[0]}
}

func (t *Tree) Root() *TreeNode {
	return t.root
}

func (t *Tree) Empty() bool {
	return t.root == nil
}

func (t *Tree) PreOrder() []any {
	ret := []any{}

	PreOrder(t.root, &ret)

	return ret
}

func (t *Tree) InOrder() []any {
	ret := []any{}

	InOrder(t.root, &ret)

	return ret
}

func (t *Tree) PostOrder() []any {
	ret := []any{}

	PostOrder(t.root, &ret)

	return ret
}

func PreOrder(root *TreeNode, ret *[]any) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)
	PreOrder(root.Left, ret)
	PreOrder(root.Right, ret)
}

func InOrder(root *TreeNode, ret *[]any) {
	if root == nil {
		return
	}

	InOrder(root.Left, ret)
	*ret = append(*ret, root.Val)
	InOrder(root.Right, ret)
}

func PostOrder(root *TreeNode, ret *[]any) {
	if root == nil {
		return
	}

	PostOrder(root.Left, ret)
	PostOrder(root.Right, ret)
	*ret = append(*ret, root.Val)
}

// func (t *Tree) InOrder() []any {
// 	var ret []any

// 	for t.root != nil {
// 		t.root = t.root.Left
// 		ret = append(ret, t.root.Val)
// 		t.root = t.root.Right
// 	}

// 	return ret
// }

// func (t *Tree) PostOrder() []any {
// 	var ret []any

// 	for t.root != nil {
// 		t.root = t.root.Left
// 		t.root = t.root.Right
// 		ret = append(ret, t.root.Val)
// 	}

// 	return ret
// }
