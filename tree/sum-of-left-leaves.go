package tree

func SumOfLeftLeavesByRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSumValue := SumOfLeftLeavesByRecursive(root.Left)

	// 如何理解, 在此处判断左叶子节点
	leftValue := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val.(int)
	}

	rightSumValue := SumOfLeftLeavesByRecursive(root.Right)

	return leftValue + leftSumValue + rightSumValue
}

func SumOfLeftLeavesByRecursiveSimple(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSumValue := SumOfLeftLeavesByRecursive(root.Left)

	// 如何理解, 在此处判断左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftSumValue = root.Left.Val.(int)
	}

	rightSumValue := SumOfLeftLeavesByRecursive(root.Right)

	return leftSumValue + rightSumValue
}

func SumOfLeftLeavesByIteration(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 左叶子节点的特性翻译为: 左节点为空，右节点为空
		if curNode.Left != nil && curNode.Left.Left == nil && curNode.Left.Right == nil {
			result += curNode.Left.Val.(int)
		}

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}

		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}

	return result
}
