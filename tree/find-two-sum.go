package tree

func FindTwoSum(root *TreeNode, target int) bool {
	hashSum := map[int]bool{}

	var find func(node *TreeNode) bool
	find = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		if _, ok := hashSum[target-root.Val.(int)]; ok {
			return true
		}

		hashSum[node.Val.(int)] = true

		return find(node.Left) || find(node.Right)
	}

	return find(root)
}

func FindTwoSumII(root *TreeNode, target int) bool {
	stack := []*TreeNode{}
	hashSum := map[int]bool{}

	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := hashSum[target-curNode.Val.(int)]; ok {
			return true
		}

		hashSum[curNode.Val.(int)] = true

		curNode = curNode.Right
	}

	return false
}

// 扩展: 是否存在相邻两个元素和为目标值

func FindTwoSumAdjacent(root *TreeNode, target int) bool {
	stack := []*TreeNode{}

	var preNode *TreeNode
	curNode := root

	for curNode != nil || len(stack) > 0 {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			curNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if preNode != nil && preNode.Val.(int)+curNode.Val.(int) == target {
				return true
			}

			preNode = curNode
			curNode = curNode.Right
		}
	}

	return false
}
