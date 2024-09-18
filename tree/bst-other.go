package tree

import "math"

//  10
//  /\
// 5 15
//   /\
//   6 20

// 错误判断模版
// func IsValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if root.Left != nil {
// 		if root.Val.(int) < root.Left.Val.(int) {
// 			return false
// 		}
// 	}

// 	if root.Right != nil {
// 		if root.Val.(int) > root.Right.Val.(int) {
// 			return false
// 		}
// 	}

// 	left := IsValidBST(root.Left)
// 	right := IsValidBST(root.Right)

// 	return left && right
// }

var (
	MAX_VAL = math.MinInt
)

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := IsValidBST(root.Left)
	if root.Val.(int) > MAX_VAL {
		MAX_VAL = root.Val.(int)
	} else {
		return false
	}

	right := IsValidBST(root.Right)

	return left && right
}

// 中序遍历, 前一个节点 < 当前节点
func IsValidBSTByIteration(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{}
	var preNode *TreeNode

	curNode := root

	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if preNode != nil && curNode.Val.(int) < preNode.Val.(int) {
			return false
		}

		preNode = curNode
		curNode = curNode.Right
	}

	return true
}

// 非递减二叉搜索树
// [1,2,3,4,5,5,6,7,8]
func FindMode(root *TreeNode) []any {
	result := make([]any, 0)

	var preNode *TreeNode
	var traversal func(node *TreeNode)

	count := 0
	maxCount := 0

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)

		// if preNode == nil {
		// 	count = 1
		// } else if preNode.Val.(int) == node.Val.(int) {
		// 	count++
		// } else {
		// 	count = 1
		// }
		// 合并分支优化
		if preNode != nil && preNode.Val.(int) == node.Val.(int) {
			count++
		} else {
			count = 1
		}

		preNode = node
		if count == maxCount {
			result = append(result, node.Val)
		}

		if count > maxCount {
			result = result[:0]
			result = []any{node.Val}
			maxCount = count
		}

		traversal(node.Right)
	}

	traversal(root)

	return result
}

func FindModeByIteration(root *TreeNode) any {
	var preNode *TreeNode
	result := make([]any, 0)

	count, maxCount := 0, 0
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// if preNode == nil {
		// 	count = 1
		// } else if preNode != nil && curNode.Val.(int) == preNode.Val.(int) {
		// 	count++
		// } else {
		// 	count = 1
		// }
		// 合并优化
		if preNode != nil && preNode.Val.(int) == curNode.Val.(int) {
			count++
		} else {
			count = 1
		}

		preNode = curNode
		if count == maxCount {
			result = append(result, curNode.Val)
		}

		if count > maxCount {
			result = []any{curNode.Val}
			maxCount = count
		}

		curNode = curNode.Right
	}

	return result
}

// 1) root.Val < low
// 2) low <= root.Val <= high
// 3) root.Val > high
// 只保留node节点值在[low, high]之间的TreeNode
func TrimBSTreeByRecursive(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val.(int) < low {
		return TrimBSTreeByRecursive(root.Right, low, high)
	}

	if root.Val.(int) > high {
		return TrimBSTreeByRecursive(root.Left, low, high)
	}

	root.Left = TrimBSTreeByRecursive(root.Left, low, high)
	root.Right = TrimBSTreeByRecursive(root.Right, low, high)

	return root
}

func TrimBSTreeByIteration(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1) root.Val < low root.Val > high
	for root != nil && (root.Val.(int) < low || root.Val.(int) > high) {
		if root.Val.(int) < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	// 2) low <= root.Val <= high
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val.(int) < low {
			cur.Left = cur.Left.Right
		}

		cur = cur.Left
	}

	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val.(int) > high {
			cur.Right = cur.Right.Left
		}

		cur = cur.Right
	}

	return root
}
