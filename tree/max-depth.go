package tree

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

func maxDepthHelper(root *TreeNode, depth int, result *int) {
	if depth > *result {
		*result = depth
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		depth++
		maxDepthHelper(root.Left, depth, result)
		depth--
	}

	if root.Right != nil {
		depth++
		maxDepthHelper(root.Right, depth, result)
		depth--
	}
}
func MaxDepthByBackTracing(root *TreeNode) int {
	result := 0
	if root == nil {
		return 0
	}

	maxDepthHelper(root, 1, &result)

	return result
}

func MaxDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}

			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}

		depth++
	}

	return depth
}
