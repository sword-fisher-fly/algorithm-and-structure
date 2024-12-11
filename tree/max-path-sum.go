package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(*TreeNode) int
	dfs = func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		leftGain := max(dfs(t.Left), 0)
		rightGain := max(dfs(t.Right), 0)

		maxSum = max(maxSum, t.Val.(int)+leftGain+rightGain)

		return t.Val.(int) + max(leftGain, rightGain)
	}

	dfs(root)
	return maxSum
}
