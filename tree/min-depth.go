package tree

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// left := MinDepth(root.Left)
	// right := MinDepth(root.Right)
	// if left == 0 || right == 0 {
	// 	return left + right + 1
	// }
	if root.Left == nil && root.Right != nil {
		return 1 + MinDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + MinDepth(root.Left)
	}

	return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
}

func MinDepthII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MinDepth(root.Left)
	right := MinDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return min(left, right) + 1
}

func MinDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		for i := 0; i < levelSize; i++ {
			curNode := queue[i]

			if curNode.Left == nil && curNode.Right == nil {
				return depth
			}

			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}

			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}

	return depth
}
