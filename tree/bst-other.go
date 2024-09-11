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

func FindMode(root *TreeNode) any {

	return 0
}

func FindModeByIteration(root *TreeNode) any {

	return 0
}
