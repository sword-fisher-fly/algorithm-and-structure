package tree

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)
}

func IsSymmetricByBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root.Left, root.Right}

	for len(stack) > 0 {
		left := stack[len(stack)-2]
		right := stack[len(stack)-1]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left, right.Right)
		stack = append(stack, left.Right, right.Left)
	}

	return true
}
