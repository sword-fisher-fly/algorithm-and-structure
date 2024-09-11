package tree

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeight(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func IsBalancedTreeByRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if getHeight(root) == -1 {
		return false
	}

	return true
}

func IsBalancedTreeByIteration(root *TreeNode) bool {

	return true
}
