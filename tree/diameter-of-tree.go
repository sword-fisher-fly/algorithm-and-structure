package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func DiameterOfBinaryTree(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	var getDepth func(*TreeNode) int
	getDepth = func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		left := getDepth(t.Left)
		right := getDepth(t.Right)
		return max(left, right) + 1
	}

	res := 0

	var preOrder func(*TreeNode)
	preOrder = func(t *TreeNode) {
		if t != nil {
			preOrder(t.Left)
			left, right := getDepth(t.Left), getDepth(t.Right)
			fmt.Printf("left=%d, right=%d\n", left, right)
			res = max(res, left+right)
			preOrder(t.Right)
		}
	}
	preOrder(root)

	return res
}

func DiameterOfBinaryTreeII(root *TreeNode) int {
	res := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	var depth func(*TreeNode) int
	depth = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		left := depth(t.Left)
		right := depth(t.Right)
		res = max(res, left+right+1)
		return max(left, right) + 1
	}

	depth(root)

	return res - 1
}
